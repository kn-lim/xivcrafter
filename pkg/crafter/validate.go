package crafter

import (
	"fmt"
)

// required flags cannot use empty strings
var REQUIRED_FLAGS = [5]string{
	"macro1",
	"confirm",
	"cancel",
	"startPause",
	"stop",
}

var ACCEPTED_KEYS = [73]string{
	"",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"0",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"insert",
	"delete",
	"home",
	"end",
	"pageup",
	"pagedown",
	"f1",
	"f2",
	"f3",
	"f4",
	"f5",
	"f6",
	"f7",
	"f8",
	"f9",
	"f10",
	"f11",
	"f12",
	"num0",
	"num1",
	"num2",
	"num3",
	"num4",
	"num5",
	"num6",
	"num7",
	"num8",
	"num9",
	"num.",
	"num+",
	"num-",
	"num*",
	"num/",
	"num_clear",
	"num_enter",
	"num_equal",
}

// CheckKeys calls checkKey to verify user input keys as valid
func CheckKeys(xiv XIVCrafter) bool {
	SUCCESS := 0

	SUCCESS += checkKey("food", xiv.Food)
	SUCCESS += checkKey("potion", xiv.Potion)
	SUCCESS += checkKey("macro1", xiv.Macro1)
	SUCCESS += checkKey("macro2", xiv.Macro2)
	SUCCESS += checkKey("confirm", xiv.Confirm)
	SUCCESS += checkKey("cancel", xiv.Cancel)
	SUCCESS += checkKey("startPause", xiv.StartPause)
	SUCCESS += checkKey("stop", xiv.Stop)

	return SUCCESS == 0
}

// checkKey validates user input keys
func checkKey(flag string, key string) int {
	for _, KEY := range ACCEPTED_KEYS {
		if KEY == key {
			for _, FLAG := range REQUIRED_FLAGS {
				if FLAG == flag {
					// check for empty string
					if key == "" {
						s := fmt.Sprintf("ERROR: FLAG %s SET WITH NO KEY", flag)
						fmt.Println(s)
						return 1
					}

					return 0
				}
			}

			return 0
		}
	}

	s := fmt.Sprintf("ERROR: FLAG %s SET WITH INVALID KEY %s", flag, key)
	fmt.Println(s)
	return 1
}

// CheckFlags calls checkFlag to verify flag values as valid
func CheckFlags(xiv XIVCrafter) bool {
	SUCCESS := 0

	// check Macro1
	SUCCESS += checkFlag("macro1Duration", xiv.Macro1Duration)

	// check Macro2
	if xiv.Macro2 != "" {
		SUCCESS += checkFlag("macro2Duration", xiv.Macro2Duration)
	}

	// check FoodDuration
	if xiv.Food != "" {
		SUCCESS += checkFlag("foodDuration", xiv.FoodDuration)
	}

	return SUCCESS == 0
}

// checkFlag validates flag value
func checkFlag(flag string, duration int) int {
	if flag == "macro1Duration" || flag == "macro2Duration" {
		if duration <= 0 {
			s := fmt.Sprintf("ERROR: FLAG %s SET WITH INVALID DURATION %d", flag, duration)
			fmt.Println(s)
			return 1
		}
	} else if flag == "foodDuration" {
		switch duration {
		case 1800:
		case 2400:
		case 2700:
		default:
			s := fmt.Sprintf("ERROR: FLAG %s SET WITH INVALID DURATION %d", flag, (duration / 60))
			fmt.Println(s)
			return 1
		}
	}

	return 0
}
