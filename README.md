# XIVCrafter
![Go](https://img.shields.io/github/go-mod/go-version/kn-lim/xivcrafter)
[![Release](https://img.shields.io/github/v/release/kn-lim/xivcrafter)](https://github.com/kn-lim/xivcrafter/releases)
![Build](https://github.com/kn-lim/xivcrafter/actions/workflows/build.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kn-lim/xivcrafter)](https://goreportcard.com/report/github.com/kn-lim/xivcrafter)
![License](https://img.shields.io/github/license/kn-lim/xivcrafter)

Automatically activates multiple crafting macros while refreshing food and potion buffs.

Tested on Windows and Keyboard only.

## Packages

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [robotgo](https://github.com/go-vgo/robotgo)
- [gohook](https://github.com/robotn/gohook)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)

# Using the Tool

Download the Windows binary in the [Releases](https://github.com/kn-lim/xivcrafter/releases) page.

## How to Run

Run:
```yml
.\xivcrafter
```

If `.xivcrafter.json` does not exist in your home directory, XIVCrafter will create it upon launch.

## Flags

Only needed if you want different settings than `.xivcrafter.json`.

```
      --cancel string        cancel hotkey
  -c, --config string        config file (default is $HOME/.xivcrafter.json)
      --confirm string       confirm hotkey
  -h, --help                 help for xivcrafter
      --start-pause string   start/pause xivcrafter hotkey
      --stop string          stop xivcrafter hotkey
```

## Accepted Keys

https://github.com/vcaesar/keycode/blob/main/keycode.go

## Prepping the Game

In order for XIVCrafter to work properly:

1. **Make sure you are not near anything that can be interacted with.**
    - This is important to make sure you don't accidentally target something else and thus being unable to craft.
2. **Open the Crafting Log and select the item you want to craft with XIVCrafter.**
3. **To ensure your character is in the correct state, start and then cancel the craft without any additional inputs.**

Once that is done, press the _Start/Pause XIVCrafter_ hotkey to start the tool.

## Config

Default Location: `$HOME/.xivcrafter.json`

Config File:

```json
{
  "start_pause": "",
  "stop": "",
  "confirm": "",
  "cancel": "",
  "recipes": [
    {
      "name": "",
      "food": "",
      "food_duration": 30,
      "potion": "",
      "macro1": "",
      "macro1_duration": 1,
      "macro2": "",
      "macro2_duration": 1,
      "macro3": "",
      "macro3_duration": 1
    }
  ]
}
```

- `start_pause`: XIVCrafter hotkey to **start or pause the program**
- `stop`: XIVCrafter hotkey to **stop the program**
- `confirm`: FFXIV hotkey for the **confirm** action
- `cancel`: FFXIV hotkey for the **cancel** action
- `recipes`: JSON list of settings representing a crafting recipe in FFXIV
  - `name`: Name of the crafting recipe
  - `food`: FFXIV hotkey for the **food** item to use
    - Leave as `""` if no food is needed
  - `food_duration`: How long the food will last (minutes)
    - Default: `30` minutes
    - Should be either 30, 40 or 45 minutes depending on whether food buffs are used
  - `potion`: FFXIV hotkey for the **potion** item to use
    - Leave as `""` if no potion is needed
    - The default length of the potion is `15` minutes
  - `macro1`: FFXIV hotkey for the **first crafting macro**
    - This field must have a value for the recipe to be valid
  - `macro1_duration`: Duration the **first crafting macro** (seconds)
  - `macro2`: FFXIV hotkey for the **second crafting macro**
    - Leave as `""` if no second crafting macro is needed
  - `macro2_duration`: Duration the **second crafting macro** (seconds)
  - `macro3`: FFXIV hotkey for the **third crafting macro**
    - Leave as `""` if no third crafting macro is needed
  - `macro3_duration`: Duration the **third crafting macro** (seconds)

# FAQ

- **Does the game need to be in focus?**
  - Yes. Otherwise, whatever program is in focus will receive the inputs.
- **Am I able to use the keyboard to type/move while the program is active?**
  - No, since XIVCrafter tracks all key presses and may act accordingly to the config provided.
- **Am I able to use the mouse while the program is active?**
  - No, as it may cause XIVCrafter to malfunction and not start the craft properly.
- **How do I get the macro duration?**
  - Count all the seconds the macro steps delays for.
  - General Rule: # of Lines * 3
- **Will this work with any craft?**
  - As long as you are able to start the craft, the program will work on any craft.
- **My craft didn't complete! What happened?**
  - Usually, latency can prevent the keys from being inputted properly to the client. If this happens, cancel the craft and wait till the program completes the "craft". Make sure to stop any existing crafting macro before the program starts a new craft. It should continue without having to reapply food and potion buffs.
