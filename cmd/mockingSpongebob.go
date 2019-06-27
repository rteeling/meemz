/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/rteeling/meemz/imgflip"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

var outputPath string

// mockingSpongebobCmd represents the mockingSpongebob command
var mockingSpongebobCmd = &cobra.Command{
	Use:   `mockingSpongebob "top text" "bottom text" path/to/image`,
	Short: "Create a Mocking Spongebob Meme",
	Long: `Create a Mocking Spongebob Meme
Usage:

meemz mockingSpongebob "text above" "textBelow" [path/to/image]
ex.
meemz mockingSpongebob "i CaN" "MaKe LoTs Of MeMes" mymeme.jpg
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for username and password
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		password := string(bytePassword)
		fmt.Print("\n\n\n")

		memeURL := imgflip.CreateMeme("102156234", args[0], args[1], strings.TrimSpace(username), strings.TrimSpace(password))
		imgflip.DownloadFile(outputPath, memeURL)
	},
}

func init() {
	rootCmd.AddCommand(mockingSpongebobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mockingSpongebobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mockingSpongebobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mockingSpongebobCmd.Flags().StringVarP(&outputPath, "output", "o", "MockingSpongebob.jpg", "Output file location")
}
