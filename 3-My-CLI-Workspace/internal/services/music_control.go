package services

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const mpvSocket = "/tmp/my-workspace-mpv"

func getPlayerCommand(ctx context.Context, filePath string) *exec.Cmd {
	switch runtime.GOOS {
	case "darwin":
		return exec.CommandContext(ctx, "afplay", filePath)
	case "windows":
		return exec.CommandContext(ctx, "ffplay", "-nodisp", "-autoexit", filePath)
	default:
		os.Remove(mpvSocket)
		return exec.CommandContext(ctx, "mpv", "--no-video", "--quiet", "--input-ipc-server="+mpvSocket, filePath)
		// return exec.CommandContext(ctx, "mpv", "--no-video", "--quiet", filePath)
	}
}

func sendSeekCommand(target string) {
	if runtime.GOOS != "linux" {
		return
	}
	conn, err := net.Dial("unix", mpvSocket)
	if err == nil {
		defer conn.Close()
		fmt.Fprintf(conn, "seek %s\n", target)
	}
}

func sendPauseToggleCommand() {
	if runtime.GOOS != "linux" {
		return
	}
	conn, err := net.Dial("unix", mpvSocket)
	if err == nil {
		defer conn.Close()
		fmt.Fprint(conn, "cycle pause\n")
	}
}

func ProcessInput(input string) (action string, value int) {
	switch input {
	case "", "n", "next":
		return "next", 0
	case "p", "prev", "b", "back":
		return "prev", 0
	case "q", "quit", "stop":
		return "quit", 0
	case "l":
		return "seek_rel", 5
	case "j":
		return "seek_rel", -5
	case "k", "pause", "play":
		return "toggle_pause", 0
	default:
		parts := strings.Split(input, ":")
		if len(parts) == 2 {
			m, err1 := strconv.Atoi(parts[0])
			s, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				return "seek_abs", m * 60 + s
			}
		}
		return "none", 0
	}
}