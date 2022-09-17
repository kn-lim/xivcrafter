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
- [progressbar](https://github.com/schollz/progressbar)

# Using the Tool

**Download the Windows 64-bit binary in the [Releases](https://github.com/kn-lim/xivcrafter/releases) page.**

## How to Build

Make sure to follow the [robotgo installation instructions](https://github.com/go-vgo/robotgo#requirements) in order to build.

Run:
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
  config      Prints and validates XIVCrafter's configuration
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

### Example Command

```
.\xivcrafter.exe start -c config.yaml --amount=10 --food="f" --potion="" --macro2=""
```

The above command will:
- Start XIVCrafter
- Use `config.yaml` as its configuration file
- Set the amount to craft to 10, rather than the value in the config file
- Sets the food hotkey to "**f**", rather than the value in the config file
- Sets the potion hotkey to "" (Empty string indicates no hotkey), rather than the value in the config file

## Supported Keys

https://github.com/kn-lim/xivcrafter/blob/main/pkg/crafter/README.md#accepted-keys

## Prepping the Game

In order for XIVCrafter to work properly:

1. **Make sure you are not near anything that can be interacted with.**
    - This is important to make sure you don't accidentally target something else and thus being unable to craft.
2. **Open the Crafting Log and select the item you want to craft with XIVCrafter.**
    - To ensure your character is in the correct state, start and then cancel the craft without any additional inputs.

Once that is done, press the _Start/Pause XIVCrafter_ hotkey to start the tool.

## User Interface

The progress bar indicates the crafting progress. It looks like the following:
```
STATUS  CURR_% [======>                       ] (CURR_AMOUNT/MAX_AMOUNT) [TIME_ELAPSED:ESTIMATED_TIME]
```
- **STATUS**: Current status of XIVCrafter.
  - **Crafting**: Currently crafting
  - **Pausing**: Will pause as soon as the current craft is completed
  - **Paused**: Currently paused and waiting to be started
  - **Stopping**: Will stop as soon as the current craft is completed
- **CURR_%**: Percentage of completion
- **CURR_AMOUNT**: Current amount crafted
- **MAX_AMOUNT**: Maximum amount to be crafted
- **TIME_ELAPSED**: How much time elapsed while crafting
- **ESTIMATED_TIME**: Estimate of how much time till completion

**NOTE**: Using verbose mode `-v` will hide the progress bar.

# FAQ

- **Does the game need to be in focus?**
  - Yes. Otherwise, whatever program is in focus will receive the inputs.
- **Am I able to use the keyboard to type/move while the program is active?**
  - No, since XIVCrafter tracks all key presses and may act accordingly to the flags provided.
- **Am I able to use the mouse while the program is active?**
  - No, as it may cause the crafter to malfunction and not start the craft properly.
- **How do I get the macro duration?**
  - Count all the seconds the macro steps delays for.
  - General Rule: # of Lines * 3
- **Will this work with any craft?**
  - As long as you are able to start the craft, the program will work on any craft.
- **My craft didn't complete! What happened?**
  - Usually, latency can prevent the keys from being inputted properly to the client. If this happens, cancel the craft and wait till the program completes the "craft". It should continue without having to reapply food and potion buffs.

# TODO

- Game no longer needs to be in focus
