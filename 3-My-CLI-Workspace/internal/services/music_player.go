package services

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

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
		fmt.Println("\n🔴 Stopping music player...")
		cancel()
	}()

	controlChan := startKeyboardListener()

	color.Green("🎵 Starting Music Player...")
	color.Yellow("💻 Controls: [Enter]/n=Next | p=Prev | l=+5s | j=-5s | 01:25=Seek | q=Quit")
	fmt.Println(strings.Repeat("-", 65))

	i := 0
	for i >= 0 && i < len(playlist) {
		if ctx.Err() != nil { break }

		track := playlist[i]
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

		elapsed := 0
		ticker := time.NewTicker(1 * time.Second)
		fmt.Println()
		fmt.Print("💻 Control: ")
		printStatus := func() {
			m := elapsed / 60
			s := elapsed % 60
			fmt.Printf("\033[s\033[1A\r\033[K%s - %02d:%02d\033[u", color.CyanString("▶ Playing (%d/%d): %s", i+1, len(playlist), filepath.Base(track)), m, s)
		}
		printStatus()

		action := "next"
		trackFinished := false

	WaitLoop:
		for {
			select {
			case <-ctx.Done():
				action = "quit"
				break WaitLoop
			case <-trackDone:
				trackFinished = true
				action = "next"
				break WaitLoop
			case <-ticker.C:
				elapsed++
				printStatus()
			case input := <-controlChan:
				fmt.Print("\033[1A\r\033[K💻 Control: ")
				act, val := ProcessInput(input)
				if act == "next" || act == "prev" || act == "quit" {
					action = act
					break WaitLoop
				} else if act == "seek_rel" {
					elapsed += val
					if elapsed < 0 { elapsed = 0 }
					sendSeekCommand(fmt.Sprintf("%d relative", val))
					printStatus()
				} else if act == "seek_abs" {
					elapsed = val
					sendSeekCommand(fmt.Sprintf("%d absolute", val))
					printStatus()
				}
			}

		}

		ticker.Stop()
		fmt.Print("\r\033[K")
		trackCancel()
		if !trackFinished { <-trackDone }
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

	color.Green("\n🟢 Playback session ended.\n")
}