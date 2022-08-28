package crafter

import (
	"fmt"
)

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

func checkKey(flag string, key string) int {
	// required flags cannot use empty strings as keys
	REQUIRED_FLAGS := [5]string{
		"macro1",
		"confirm",
		"cancel",
		"startPause",
		"stop",
	}

	ACCEPTED_KEYS := [55]string{
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
	}

	for _, KEY := range ACCEPTED_KEYS {
		if KEY == key {
			for _, FLAG := range REQUIRED_FLAGS {
				if FLAG == flag {
					// check for empty string
					if key == "" {
						s := fmt.Sprintf("ERROR: FLAG %s SET WITH NO KEY", flag)
						fmt.Println(s)
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
