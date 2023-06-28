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
- [lipgloss](https://github.com/charmbracelet/lipgloss)

# Using the Tool

**Download the Windows binary in the [Releases](https://github.com/kn-lim/xivcrafter/releases) page.**

## How to Run

Run:
```yml
.\xivcrafter
```

If `.xivcrafter.json` does not exist in your home directory, running XIVCrafter will create it upon launch.

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
