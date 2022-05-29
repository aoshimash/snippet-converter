/*
Copyright © 2022 Shuji Aoshima <aoshimash@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/aoshimash/snippet-converter/internal/codesnippet"
	"github.com/aoshimash/snippet-converter/internal/util"
	"github.com/spf13/cobra"
)

// genvscodesnippetsCmd represents the genvscodesnippets command
var genvscodesnippetsCmd = &cobra.Command{
	Use:   "genvscodesnippets [PATH]",
	Short: "Generate a JSON string for Visual Studio Code Snippets",
	Long:  `Generate a JSON string for Visual Studio Code Snippets from comments in source codes.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		genvscodesnippets(args[0])
	},
}

func init() {
	rootCmd.AddCommand(genvscodesnippetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genvscodesnippetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genvscodesnippetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genvscodesnippets(targetDir string) {
	// get all file pathes
	filePathes, err := util.GetFilePathes(targetDir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// create slice of CodeSnipet objects from file data
	snippets := codesnippet.NewCodeSnippets(filePathes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// get json
	bytes, err := codesnippet.GetVSCodeSnippetsJSON(snippets)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(bytes))
}
