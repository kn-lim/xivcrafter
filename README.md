# XIVCrafter
![Go](https://img.shields.io/github/go-mod/go-version/kn-lim/xivcrafter)
[![Release](https://img.shields.io/github/v/release/kn-lim/xivcrafter)](https://github.com/kn-lim/xivcrafter/releases)
![Build](https://img.shields.io/github/actions/workflow/status/kn-lim/xivcrafter/release.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kn-lim/xivcrafter)](https://goreportcard.com/report/github.com/kn-lim/xivcrafter)
![License](https://img.shields.io/github/license/kn-lim/xivcrafter)

A FFXIV CLI app that activates multiple crafting macros while refreshing food and potion buffs.

Tested on Windows and Keyboard only.

<p align="center">
  <img src="https://raw.githubusercontent.com/kn-lim/kn-lim/main/xivcrafter/xivcrafter_ui.gif"></img>
</p>

## Packages Used

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [robotgo](https://github.com/go-vgo/robotgo)
- [gohook](https://github.com/robotn/gohook)

# Using the Tool

Download the Windows binary in the [Releases](https://github.com/kn-lim/xivcrafter/releases) page.

**Need help? Check out the [Wiki](https://github.com/kn-lim/xivcrafter/wiki/).**

## How to Run

Run:
```yml
.\xivcrafter
```

If `.xivcrafter.json` does not exist in your home directory, XIVCrafter will create it upon launch. It will then prompt you to fill out the XIVCrafter settings and one recipe.

## Flags

Only needed if you want different settings than `.xivcrafter.json`.

```
      --cancel string        cancel hotkey
  -c, --config string        config file (default is $HOME/.xivcrafter.json)
      --confirm string       confirm hotkey
      --debug                enable debugging (debug log location is $HOME/.xivcrafter-debug.log)
  -h, --help                 help for xivcrafter
      --start-pause string   start/pause xivcrafter hotkey
      --stop string          stop xivcrafter hotkey
```

## Prepping the Game

In order for XIVCrafter to work properly:

1. **Make sure you are not near anything that can be interacted with.**
    - This is important to make sure you don't accidentally target something else and thus being unable to craft.
2. **Open the Crafting Log and select the item you want to craft with XIVCrafter.**
3. **To ensure your character is in the correct state, start and then cancel the craft without any additional inputs.**
    - Press the **Synthesize** button in the Crafting Log window. When the Synthesize window appears, press the **Quit** button.

Once that is done, press the **Start/Pause XIVCrafter** hotkey to start XIVCrafter.

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
- `recipes`: JSON array of settings representing a crafting recipe in FFXIV
  - `name`: Name of the crafting recipe
  - `food`: _OPTIONAL_ FFXIV hotkey for the **food** item to use
    - Leave as `""` if no food is needed
  - `food_duration`: How long the food will last (minutes)
    - Default: `30` minutes
    - Should be either 30, 40 or 45 minutes depending on whether food buffs are used
  - `potion`: _OPTIONAL_ FFXIV hotkey for the **potion** item to use
    - Leave as `""` if no potion is needed
    - The default length of the potion is `15` minutes
  - `macro1`: FFXIV hotkey for the **first crafting macro**
    - This field must have a value for the recipe to be valid
  - `macro1_duration`: Duration the **first crafting macro** (seconds)
  - `macro2`: _OPTIONAL_ FFXIV hotkey for the **second crafting macro**
    - Leave as `""` if no second crafting macro is needed
  - `macro2_duration`: Duration the **second crafting macro** (seconds)
  - `macro3`: _OPTIONAL_ FFXIV hotkey for the **third crafting macro**
    - Leave as `""` if no third crafting macro is needed
  - `macro3_duration`: Duration the **third crafting macro** (seconds)

For more information on filling out the config file, go to the [Config](https://github.com/kn-lim/xivcrafter/wiki/User-Guide#config) section of the User Guide.

## Accepted Keys

https://github.com/kn-lim/xivcrafter/wiki/User-Guide#accepted-keys

# FAQ

<details>
<summary>
<b>Does the game need to be in focus?</b>
</summary>
<p>Yes. Otherwise, whatever program is in focus will receive the inputs.</p>
</details>

<details>
<summary>
<b>Am I able to use the keyboard to type/move while the program is active?</b>
</summary>
<p>No, since XIVCrafter tracks all key presses and may act accordingly to the config provided.</p>
</details>

<details>
<summary>
<b>Am I able to use the mouse while the program is active?</b>
</summary>
<p>No, as it may cause XIVCrafter to malfunction and not start the craft properly.</p>
</details>

<details>
<summary>
<b>How do I get the macro duration?</b>
</summary>
<p>Count all the seconds the macro steps delays for.</p>
<p>General Rule: # of Lines * 3</p>
</details>

<details>
<summary>
<b>Will this work with the latest patch?</b>
</summary>
<p>Unless the <a href="https://github.com/go-vgo/robotgo">robotgo</a> package stops working or FFXIV blocks virtual keyboard inputs, XIVCrafter should work on any patch.</p>
</details>

<details>
<summary>
<b>Will this work with any craft?</b>
</summary>
<p>As long as you are able to start the craft, XIVCrafter will work on any craft.</p>
</details>

<details>
<summary>
<b>My craft didn't complete! What happened?</b>
</summary>
<p>Usually, latency can prevent the keys from being inputted properly to the client. If this happens, cancel the craft manually in-game. Get back into the <a href="#prepping-the-game">initial starting state</a> and wait till XIVCrafter completes the "craft". Make sure to stop any existing crafting macro before XIVCrafter starts a new craft. It should continue without having to reapply food and potion buffs.</p>
<p>You may need to cancel the current active crafting macro in order to get back to the initial starting state. To do that, you will need to interrupt that macro. You can do that by having this as a macro: <code>/e end</code>. By activating that one line macro, it should interrupt any currently running crafting macro to allow you to get back into the initial starting state. </p>
</details>

<br>For more information, go to the [Wiki](https://github.com/kn-lim/xivcrafter/wiki/).
