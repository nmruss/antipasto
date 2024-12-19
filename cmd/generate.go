package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates a valid project at the specified folder",
	Long:    `Project generation will create a new, valid project folder structure at the root of a specified folder`,
	Aliases: []string{"g"},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErrf("No arguments found. Please provide an empty folder path as the first argument\n")
			return
		}

		createFoldersAtPath(args[0])
	},
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createFoldersAtPath(rootPath string) {
	folderExists, err := exists(rootPath)
	mkDirErr := os.MkdirAll(rootPath, 0755)

	if mkDirErr != nil {
		panic(mkDirErr)
	}

	if err != nil {
		panic(err)
	}

	if folderExists {
		fmt.Println("Folder already exists, please pass a new path for project generation")
	} else {
		fmt.Printf("New project folder created at %s \n", rootPath)
	}

}
