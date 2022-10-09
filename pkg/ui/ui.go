package ui

import (
	"github.com/schollz/progressbar/v3"
)

// Constants

const START = "[green]Crafting[reset]"
const PAUSE = "[red]Pausing [reset]"
const STOP = "[red]Paused  [reset]"
const EXIT = "[red]Stopping[reset]"
const WIDTH = 30

// UI Struct
type UI struct {
	PB *progressbar.ProgressBar
}

// Init initializes the UI struct
func (ui *UI) Init(max int) {
	PB := progressbar.NewOptions(max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(WIDTH),
		progressbar.OptionSetDescription(STOP),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	*ui = UI{
		PB: PB,
	}
}

// Increment increments the progress bar
func (ui *UI) Increment() {
	ui.PB.Add(1)
}

// SetStart sets the description of the progress bar to START
func (ui *UI) SetStart() {
	ui.PB.Describe(START)
}

// SetPause sets the description of the progress bar to PAUSE
func (ui *UI) SetPause() {
	ui.PB.Describe(PAUSE)
}

// SetStop sets the description of the progress bar to STOP
func (ui *UI) SetStop() {
	ui.PB.Describe(STOP)
}

// SetExit sets the description of the progress bar to EXIT
func (ui *UI) SetExit() {
	ui.PB.Describe(EXIT)
}

// Start renders the blank progress bar
func (ui *UI) Start() {
	ui.PB.RenderBlank()
}
