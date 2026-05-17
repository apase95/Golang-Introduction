package services

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func PlayMusic(path string, shuffle bool) {
	playlist, err := loadPlaylist(path, shuffle)
	if err != nil { color.Red("✖ %v\n", err); return }
	if len(playlist) == 0 {
		color.Yellow("⚠ No MP3 files found in: %s\n", path)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n🛑 Stopping music player...")
		cancel()
	}()

	controlChan := startKeyboardListener()

	color.Green("🎵 Starting Music Player...")
	color.Yellow("⌨  Controls: [Enter] or 'n' = Next | 'p' = Prev | 'q' = Quit")
	fmt.Println(strings.Repeat("-", 50))

	i := 0
	for i >= 0 && i < len(playlist) {
		if ctx.Err() != nil {
			break
		}

		track := playlist[i]
		color.Cyan("▶ Playing (%d/%d): %s", i + 1, len(playlist), filepath.Base(track))

		trackCtx, trackCancel := context.WithCancel(ctx)
		cmd := getPlayerCommand(trackCtx, track)
		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Start(); err != nil {
			color.Red("✖ Failed to start player: %v\n", err)
			trackCancel()
			i++
			continue
		}

		trackDone := make(chan error, 1)
		go func() {
			trackDone <- cmd.Wait()
		}()

		action, trackFinished := handleTrackControl(ctx, trackDone, controlChan)
		trackCancel()
		if !trackFinished {
			<-trackDone
		}

		if action == "quit" {
			break
		} else if action == "next" {
			i++
		} else if action == "prev" {
			if i > 0 {
				i--
			} else {
				color.Yellow("   (Already at the first track, replaying...)")
			}
		}
	}

	color.Green("\n✅ Playback session ended.\n")
}