package metrics

import (
	"github.com/andr-ii/punchy/colors"
	"github.com/andr-ii/punchy/conf"
	"github.com/andr-ii/punchy/terminal"
	"time"
)

var state_done = colors.Blue("█")
var progress_label_raw = "[ PROGRESS: "
var progress_label = colors.Blue(progress_label_raw)
var progress_row = conf.DEFAULT_HEIGHT + 2

func runProgress() {
	progressCol := len(progress_label_raw) + 1
	terminal.PrintAt(progress_row, 0, progress_label)
	terminal.PrintAt(progress_row, len(progress_label_raw)+80, colors.Blue(" ]"))

	for {
		time.Sleep(time.Duration(conf.Plan.Duration*750) * time.Millisecond)
		terminal.PrintAt(progress_row, progressCol, state_done)
		progressCol += 1
	}
}
