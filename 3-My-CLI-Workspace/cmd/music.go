package cmd

import (
	"3-My-CLI-Workspace/internal/services"

	"github.com/spf13/cobra"
)


var shuffle bool

var musicCmd = &cobra.Command{
	Use: 	"music",
	Short: 	"Local music player",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var musicPlayCmd = &cobra.Command{
	Use: "play [file_or_folder_path]",
	Short: "Play an MP3 file or all MP3s in a folder",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		services.PlayMusic(args[0], shuffle)
	},
}

func init() {
	musicPlayCmd.Flags().BoolVarP(&shuffle, "shuffle", "s", false, "Play music in random order")
	
	musicCmd.AddCommand(musicPlayCmd)
	RootCmd.AddCommand(musicCmd)
}