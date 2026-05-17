package services

import (
	"context"
	"os/exec"
	"runtime"
)

func getPlayerCommand(ctx context.Context, filePath string) *exec.Cmd {
	switch runtime.GOOS {
	case "darwin":
		return exec.CommandContext(ctx, "afplay", filePath)
	case "windows":
		return exec.CommandContext(ctx, "ffplay", "-nodisp", "-autoexit", filePath)
	default:
		return exec.CommandContext(ctx, "mpv", "--no-video", "--quiet", filePath)
	}
}

func handleTrackControl(ctx context.Context, trackDone <-chan error, controlChan <-chan string) (action string, trackFinished bool) {
	for {
		select {
		case <-ctx.Done():
			return "quit", false

		case <-trackDone:
			return "next", true

		case input := <-controlChan:
			switch input {
			case "", "n", "next":
				return "next", false
			case "p", "prev", "b", "back":
				return "prev", false
			case "q", "quit", "stop":
				return "quit", false
			default:
			}
		}
	}
}