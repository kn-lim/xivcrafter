# XIVCrafter

Automatically activates multiple crafting macros while refreshing food and potion buffs.

Works on Windows and Keyboard only.

## Requirements

| Name | Version |
|------|---------|
| go   | 1.19    |

## Packages

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [robotgo](https://github.com/go-vgo/robotgo)

# Using the Tool

Download the Windows 64-bit binary in the [Releases](https://github.com/kn-lim/xivcrafter/releases) page.

## How to Build

```yml
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build
```

## How to Run

### Using Flags

Run:
```yml
.\xivcrafter start --[FLAGS]
```

### Using a Config File

1. Modify [config.yaml](https://github.com/kn-lim/xivcrafter/blob/main/config.yaml) to match your hotkeys and crafts.
    - Default location XIVCrafter looks for the config file is `$HOME\.xivcrafter.yaml`.
2. Run:
```yml
.\xivcrafter start -c \path\to\config.yaml

# If config file is $HOME\.xivcrafter.yaml
.\xivcrafter start
```

**NOTE**: Using flags will take precendence over the config file.

## Available Commands

```
  help        Help about any command
  start       Starts XIVCrafter
```

## Flags

```
      --amount int           amount to craft
      --cancel string        cancel hotkey
  -c, --config string        config file (default is $HOME\.xivcrafter.yaml)
      --confirm string       confirm hotkey
      --food string          food hotkey
      --foodDuration int     food duration (minutes)
  -h, --help                 help for xivcrafter
      --macro1 string        macro 1 hotkey
      --macro1Duration int   macro 1 duration (seconds)
      --macro2 string        macro 2 hotkey
      --macro2Duration int   macro 2 duration (seconds)
      --potion string        potion hotkey
  -r, --random               use random delay
      --startPause string    start/pause xivcrafter hotkey
      --stop string          stop xivcrafter hotkey
  -v, --verbose              verbose output
```

## Supported Keys

https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

## Prepping the Game

In order for XIVCrafter to work properly:

1. **Make sure you are not near anything that can be interacted with.**
    - This is important to make sure you don't accidentally target something else and thus being unable to craft.
2. **Open the Crafting Log and select the item you want to craft with XIVCrafter.**
    - To ensure your character is in the correct state, start and then cancel the craft without any additional inputs.

Once that is done, press the _Start/Pause XIVCrafter_ hotkey to start the tool.

# FAQ

- **Does the game need to be in focus?**
  - Yes. Otherwise, whatever program is in focus will receive the inputs.
- **Am I able to use the keyboard to type/move while the program is active?**
  - No, since XIVCrafter tracks all key presses and may act accordingly to the flags provided.
- **How do I get the macro duration?**
  - Count all the seconds the macro steps delays for.
  - General Rule: # of Lines * 3

# TODO

- Sanitize user input properly
- Show estimated time to complete the craft
- Game no longer needs to be in focus
