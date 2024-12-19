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

		createProjectFoldersAtPath(args[0])
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

func createProjectFoldersAtPath(rootPath string) {
	rootFolderExists, rootFolderErr := exists(rootPath)
	if rootFolderErr != nil {
		panic(rootFolderErr)
	}

	if rootFolderExists {
		fmt.Println("Folder already exists, please pass a new path for project generation")
	} else {
		dirsToMake := []string{rootPath, rootPath + "/input", rootPath + "/output", rootPath + "/output/300x250", rootPath + "/output/300x250/styles", rootPath + "/output/300x250/src"}
		for _, dir := range dirsToMake {
			mkDirErr := os.MkdirAll(dir, 0755)
			if mkDirErr != nil {
				panic(mkDirErr)
			}
		}

		fmt.Printf("New project folder created at %s \n", rootPath)
	}
}
