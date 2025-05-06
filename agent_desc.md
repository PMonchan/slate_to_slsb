#### SLATEtoSLSB
SLATEActionLogs/
The json files under this directory have the following structure:
```json
{
        "stringList" :
        {
                "slate.actionlog" :
                [
                        "AddTag,Put Animation Name Here,TagToAdd",              /* Adds the "TagToAdd" tag to the animation with name "Put Animation Name Here"         */
                        "RemoveTag,Put Animation Name Here,TagToRemove",        /* Removes the "TagToRemove" tag from the animation with name "Put Animation Name Here" */
                        "Disable,Put Animation Name Here",                      /* Disables the animation with name "Put Animation Name Here"                           */
                        "AddTag,Put Animation Name Here,NoCommaOnLastLine"      /* Remember: the last line does not get a comma at the end!                             */
                ]
        }
}
```

Follow the comments in the JSON above.
SLSBjsons/ The files under this directory have the following structure:
``` json
{
  "pack_name": "Anubs Human",
  "pack_author": "Unknown",
  "prefix_hash": "e73c",
  "scenes": {
    "dl3imhri": {
      "id": "dl3imhri",
      "name": "Anubs Amulet Hypnosis",
      "stages": [
        {
          "id": "die1iedx",
          "name": "",
          "positions": [
            {
              "sex": {
                "male": false,
                "female": true,
                "futa": false
              },
              "race": "Human",
              "event": [
                "Hypno_A1_S1"
              ],
              "scale": 1.0,
              "extra": {
                "submissive": false,
                "vampire": false,
                "climax": false,
                "dead": false,
                "custom": []
              },
              "offset": {
                "x": 0.0,
                "y": 0.0,
                "z": 0.0,
                "r": 0.0
              },
              "anim_obj": "",
              "strip_data": {
                "default": true,
                "everything": false,
                "nothing": false,
                "helmet": false,
                "gloves": false,
                "boots": false
              },
              "schlong": 9
            },
            {
              "sex": {
                "male": true,
                "female": false,
                "futa": true
              },
              "race": "Human",
              "event": [
                "Hypno_A2_S1"
              ],
              "scale": 1.0,
              "extra": {
                "submissive": false,
                "vampire": false,
                "climax": false,
                "dead": false,
                "custom": []
              },
              "offset": {
                "x": 0.0,
                "y": 0.0,
                "z": 0.0,
                "r": 0.0
              },
              "anim_obj": "AnubAmulet",
              "strip_data": {
                "default": true,
                "everything": false,
                "nothing": false,
                "helmet": false,
                "gloves": false,
                "boots": false
              },
              "schlong": 5
            }
          ],
          "tags": [

          ],
          "extra": {
            "fixed_len": 0.0,
            "nav_text": ""
          }
        },
        {
          "id": "01xj80up",
          "name": "",
          "positions": [
            {
              "sex": {
                "male": false,
                "female": true,
                "futa": false
              },
              "race": "Human",
              "event": [
                "Hypno_A1_S2"
              ],
              "scale": 1.0,
              "extra": {
                "submissive": false,
                "vampire": false,
                "climax": false,
                "dead": false,
                "custom": []
              },
              "offset": {
                "x": 0.0,
                "y": 0.0,
                "z": 0.0,
                "r": 0.0
              },
              "anim_obj": "",
              "strip_data": {
                "default": true,
                "everything": false,
                "nothing": false,
                "helmet": false,
                "gloves": false,
                "boots": false
              },
              "schlong": 9
            },
            {
              "sex": {
                "male": true,
                "female": false,
                "futa": true
              },
              "race": "Human",
              "event": [
                "Hypno_A2_S2"
              ],
              "scale": 1.0,
              "extra": {
                "submissive": false,
                "vampire": false,
                "climax": false,
                "dead": false,
                "custom": []
              },
              "offset": {
                "x": 0.0,
                "y": 0.0,
                "z": 0.0,
                "r": 0.0
              },
              "anim_obj": "AnubAmulet",
              "strip_data": {
                "default": true,
                "everything": false,
                "nothing": false,
                "helmet": false,
                "gloves": false,
                "boots": false
              },
              "schlong": 6
            }
          ],
          "tags": [
          ],
          "extra": {
            "fixed_len": 0.0,
            "nav_text": ""
          }
        }
      ],
      "root": "die1iedx",
      "graph": {
        "1wcvuls0": {
          "dest": [
            "5s6e11fm"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "3thx1tjt": {
          "dest": [
            "fe1l5j4l"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "die1iedx": {
          "dest": [
            "01xj80up"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "9in8gf39": {
          "dest": [],
          "x": 40.0,
          "y": 40.0
        },
        "lhbbzbm8": {
          "dest": [
            "1wcvuls0"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "5s6e11fm": {
          "dest": [
            "3thx1tjt"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "01xj80up": {
          "dest": [
            "lhbbzbm8"
          ],
          "x": 40.0,
          "y": 40.0
        },
        "fe1l5j4l": {
          "dest": [
            "9in8gf39"
          ],
          "x": 40.0,
          "y": 40.0
        }
      },
      "furniture": {
        "furni_types": [
          "None"
        ],
        "allow_bed": false,
        "offset": {
          "x": 0.0,
          "y": 0.0,
          "z": 0.0,
          "r": 0.0
        }
      },
      "private": false,
      "has_warnings": false
    }
  }
}
```
In the SLATEActions json, "Put Animation Name Here" corresponds to scenes/id/name. Perform a search based on the above conditions and manipulate the stage/tags according to the comments in the SLATEactionlogs example.

