package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// SLATEActionLog represents the structure of a SLATE action log file
type SLATEActionLog struct {
	StringList struct {
		SlateActionLog []string `json:"slate.actionlog"`
	} `json:"stringList"`
}

// SLSBJson represents the structure of an SLSB JSON file
type SLSBJson struct {
	PackName   string               `json:"pack_name"`
	PackAuthor string               `json:"pack_author"`
	PrefixHash string               `json:"prefix_hash"`
	Scenes     map[string]SLSBScene `json:"scenes"`
}

// SLSBScene represents a scene in an SLSB JSON file
type SLSBScene struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Stages      []SLSBStage          `json:"stages"`
	Root        string               `json:"root"`
	Graph       map[string]SLSBGraph `json:"graph"`
	Furniture   SLSBFurniture        `json:"furniture"`
	Private     bool                 `json:"private"`
	HasWarnings bool                 `json:"has_warnings"`
}

// SLSBStage represents a stage in an SLSB scene
type SLSBStage struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Positions []SLSBPosition `json:"positions"`
	Tags      []string       `json:"tags"`
	Extra     struct {
		FixedLen float64 `json:"fixed_len"`
		NavText  string  `json:"nav_text"`
	} `json:"extra"`
}

// SLSBPosition represents a position in an SLSB stage
type SLSBPosition struct {
	Sex       SLSBSex       `json:"sex"`
	Race      string        `json:"race"`
	Event     []string      `json:"event"`
	Scale     float64       `json:"scale"`
	Extra     SLSBExtra     `json:"extra"`
	Offset    SLSBOffset    `json:"offset"`
	AnimObj   string        `json:"anim_obj"`
	StripData SLSBStripData `json:"strip_data"`
	Schlong   int           `json:"schlong"`
}

// SLSBSex represents the sex options for a position
type SLSBSex struct {
	Male   bool `json:"male"`
	Female bool `json:"female"`
	Futa   bool `json:"futa"`
}

// SLSBExtra represents extra options for a position
type SLSBExtra struct {
	Submissive bool     `json:"submissive"`
	Vampire    bool     `json:"vampire"`
	Climax     bool     `json:"climax"`
	Dead       bool     `json:"dead"`
	Custom     []string `json:"custom"`
}

// SLSBOffset represents the offset for a position
type SLSBOffset struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	R float64 `json:"r"`
}

// SLSBStripData represents the strip data for a position
type SLSBStripData struct {
	Default    bool `json:"default"`
	Everything bool `json:"everything"`
	Nothing    bool `json:"nothing"`
	Helmet     bool `json:"helmet"`
	Gloves     bool `json:"gloves"`
	Boots      bool `json:"boots"`
}

// SLSBGraph represents a graph node in an SLSB scene
type SLSBGraph struct {
	Dest []string `json:"dest"`
	X    float64  `json:"x"`
	Y    float64  `json:"y"`
}

// SLSBFurniture represents furniture options for an SLSB scene
type SLSBFurniture struct {
	FurniTypes []string   `json:"furni_types"`
	AllowBed   bool       `json:"allow_bed"`
	Offset     SLSBOffset `json:"offset"`
}

// Action represents a parsed action from a SLATE action log
type Action struct {
	Type          string
	AnimationName string
	Tag           string
}

// ParseAction parses a SLATE action log entry into an Action struct
func ParseAction(actionStr string) (Action, error) {
	parts := strings.Split(actionStr, ",")
	if len(parts) < 2 {
		return Action{}, fmt.Errorf("invalid action format: %s", actionStr)
	}

	action := Action{
		Type: strings.TrimSpace(strings.ToLower(parts[0])),
	}

	if len(parts) >= 2 {
		action.AnimationName = strings.TrimSpace(parts[1])
	}

	if len(parts) >= 3 && (action.Type == "addtag" || action.Type == "removetag") {
		action.Tag = strings.TrimSpace(parts[2])
	}

	return action, nil
}

// ProcessActions processes the actions from a SLATE action log and applies them to an SLSB JSON
func ProcessActions(slateActions []string, slsbJson *SLSBJson) error {
	for _, actionStr := range slateActions {
		action, err := ParseAction(actionStr)
		if err != nil {
			return err
		}

		// Find the scene with the matching name
		var sceneID string
		for id, scene := range slsbJson.Scenes {
			if scene.Name == action.AnimationName {
				sceneID = id
				break
			}
		}

		if sceneID == "" {
			continue
		}

		// Get a copy of the scene
		scene := slsbJson.Scenes[sceneID]

		// Apply the action to all stages in the scene
		switch action.Type {
		case "addtag":
			for i := range scene.Stages {
				// Check if the tag already exists
				tagExists := false
				for _, tag := range scene.Stages[i].Tags {
					if tag == action.Tag {
						tagExists = true
						break
					}
				}
				if !tagExists {
					scene.Stages[i].Tags = append(scene.Stages[i].Tags, action.Tag)
				}
			}
		case "removetag":
			for i := range scene.Stages {
				newTags := []string{}
				for _, tag := range scene.Stages[i].Tags {
					if tag != action.Tag {
						newTags = append(newTags, tag)
					}
				}
				scene.Stages[i].Tags = newTags
			}
		case "disable":
			// For "disable" action, we would need to implement the logic based on how disabling works in the SLSB format
			// This is a placeholder and would need to be updated based on the actual requirements
			fmt.Printf("Warning: 'disable' action not fully implemented for scene %s\n", sceneID)
		default:
			return fmt.Errorf("unknown action type: %s", action.Type)
		}

		// Update the scene in the map
		slsbJson.Scenes[sceneID] = scene
	}

	return nil
}

// getJsonFilesInDir returns a list of JSON files in the specified directory
func getJsonFilesInDir(dirPath string) ([]string, error) {
	var files []string
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".json") {
			files = append(files, filepath.Join(dirPath, entry.Name()))
		}
	}

	return files, nil
}

// processSlateActionLogFile processes a single SLATE action log file against all SLSB JSON files in the specified directory
func processSlateActionLogFile(slateActionLogPath string, slsbJsonDir string, outputDir string) error {
	// Read the SLATE action log file
	slateActionLogData, err := ioutil.ReadFile(slateActionLogPath)
	if err != nil {
		return fmt.Errorf("error reading SLATE action log file: %v", err)
	}

	// Parse the SLATE action log
	var slateActionLog SLATEActionLog
	err = json.Unmarshal(slateActionLogData, &slateActionLog)
	if err != nil {
		return fmt.Errorf("error parsing SLATE action log file: %v", err)
	}

	// Get all SLSB JSON files in the directory
	slsbJsonFiles, err := getJsonFilesInDir(slsbJsonDir)
	if err != nil {
		return fmt.Errorf("error reading SLSB JSON directory: %v", err)
	}

	if len(slsbJsonFiles) == 0 {
		return fmt.Errorf("no JSON files found in SLSB directory: %s", slsbJsonDir)
	}

	// Process each SLSB JSON file
	for _, slsbJsonPath := range slsbJsonFiles {
		// Determine the output path for this file
		base := filepath.Base(slsbJsonPath)
		outputPath := filepath.Join(outputDir, base)

		// Check if the file already exists in the output directory
		// If it does, read from there instead of the original directory
		var slsbJsonData []byte
		var readPath string

		if _, err := os.Stat(outputPath); err == nil {
			// File exists in output directory, read from there
			slsbJsonData, err = ioutil.ReadFile(outputPath)
			readPath = outputPath
		} else {
			// File doesn't exist in output directory, read from original directory
			slsbJsonData, err = ioutil.ReadFile(slsbJsonPath)
			readPath = slsbJsonPath
		}

		if err != nil {
			fmt.Printf("Error reading SLSB JSON file %s: %v\n", readPath, err)
			continue
		}

		// Parse the SLSB JSON
		var slsbJson SLSBJson
		err = json.Unmarshal(slsbJsonData, &slsbJson)
		if err != nil {
			fmt.Printf("Error parsing SLSB JSON file %s: %v\n", readPath, err)
			continue
		}

		// Process the actions
		err = ProcessActions(slateActionLog.StringList.SlateActionLog, &slsbJson)
		if err != nil {
			// If scene not found, just continue to the next file
			if strings.Contains(err.Error(), "scene not found") {
				continue
			}
			fmt.Printf("Error processing actions for SLSB JSON file %s: %v\n", slsbJsonPath, err)
			continue
		}

		// Write the modified SLSB JSON back to disk
		modifiedSlsbJsonData, err := json.MarshalIndent(slsbJson, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling modified SLSB JSON for file %s: %v\n", slsbJsonPath, err)
			continue
		}

		// Write the modified SLSB JSON to the output directory
		err = ioutil.WriteFile(outputPath, modifiedSlsbJsonData, 0644)
		if err != nil {
			fmt.Printf("Error writing modified SLSB JSON to file %s: %v\n", outputPath, err)
			continue
		}

		fmt.Printf("Successfully processed %d actions from %s and wrote the modified SLSB JSON to %s\n",
			len(slateActionLog.StringList.SlateActionLog),
			filepath.Base(slateActionLogPath),
			outputPath)
	}

	return nil
}

func main() {
	// Parse command-line arguments
	slateActionLogDir := flag.String("slate-dir", "", "Path to the directory containing SLATE action log files")
	slsbJsonDir := flag.String("slsb-dir", "", "Path to the directory containing SLSB JSON files")
	outputDir := flag.String("output-dir", "output", "Path to the directory where modified SLSB JSON files will be written")
	flag.Parse()

	// Check if required arguments are provided
	if *slateActionLogDir == "" || *slsbJsonDir == "" {
		fmt.Println("Usage: slateToSlsb -slate-dir <path_to_slate_action_log_directory> -slsb-dir <path_to_slsb_json_directory> [-output-dir <path_to_output_directory>]")
		os.Exit(1)
	}

	// Create the output directory if it doesn't exist
	err := os.MkdirAll(*outputDir, 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Get all SLATE action log files in the directory
	slateActionLogFiles, err := getJsonFilesInDir(*slateActionLogDir)
	if err != nil {
		fmt.Printf("Error reading SLATE action log directory: %v\n", err)
		os.Exit(1)
	}

	if len(slateActionLogFiles) == 0 {
		fmt.Printf("No JSON files found in SLATE action log directory: %s\n", *slateActionLogDir)
		os.Exit(1)
	}

	// Process each SLATE action log file
	for _, slateActionLogPath := range slateActionLogFiles {
		fmt.Printf("Processing SLATE action log file: %s\n", slateActionLogPath)
		err := processSlateActionLogFile(slateActionLogPath, *slsbJsonDir, *outputDir)
		if err != nil {
			fmt.Printf("Error processing SLATE action log file %s: %v\n", slateActionLogPath, err)
		}
	}

	fmt.Println("All SLATE action log files processed successfully!")
}
