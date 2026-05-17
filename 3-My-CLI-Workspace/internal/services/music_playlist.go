package services

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func loadPlaylist(path string, shuffle bool) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("error reading path: %v", err)
	}

	var playlist []string

	if info.IsDir() {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".mp3") {
				playlist = append(playlist, filepath.Join(path, entry.Name()))
			}
		}
	} else {
		if strings.HasSuffix(strings.ToLower(info.Name()), ".mp3") {
			playlist = append(playlist, path)
		}
	}

	if shuffle && len(playlist) > 0 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(playlist), func(i, j int) {
			playlist[i], playlist[j] = playlist[j], playlist[i]
		})
	}

	return playlist, nil
}