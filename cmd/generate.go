package cmd

import (
	"errors"
	"fmt"
	"nmruss/antipasto/configuration"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("config", "default.apconfig", "Specifies an .apconfig configuration file for project generation")
	//NOTE: add flags for asset size, and a way to mark this in the configuration file
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
		configFileName, _ := cmd.LocalFlags().GetString("config")

		config := configuration.ParseConfigurationFile(configFileName)
		createProjectFoldersAtPath(args[0], config)
		writeProjectFilesAtPath(args[0], config)
	},
}

func createProjectFoldersAtPath(rootPath string, config configuration.APConfig) error {
	rootFolderExists, rootFolderErr := exists(rootPath)
	if rootFolderErr != nil {
		panic(rootFolderErr)
	}

	if rootFolderExists {
		fmt.Println("Folder already exists, please pass a new path for project generation")
	} else {
		widthStr := strconv.Itoa(config.Size[0])
		heightStr := strconv.Itoa(config.Size[1])
		sizeStr := widthStr + "x" + heightStr
		dirsToMake := []string{rootPath, rootPath + "/input", rootPath + "/output", rootPath + "/output/" + sizeStr, rootPath + "/output/" + sizeStr + "/styles", rootPath + "/output/" + sizeStr + "/src"}
		for _, dir := range dirsToMake {
			mkDirErr := os.MkdirAll(dir, 0755)
			if mkDirErr != nil {
				panic(mkDirErr)
			}
		}

		fmt.Printf("New project folder created at %s \n", rootPath)
	}
	return nil
}

// Write project files from a configuration file to folders at rootPath
func writeProjectFilesAtPath(rootPath string, config configuration.APConfig) error {
	rootFolderExists, rootFolderErr := exists(rootPath)
	if rootFolderErr != nil {
		panic(rootFolderErr)
	}
	if !rootFolderExists {
		return errors.New("specified project folder does not exist")
	}

	if len(config.DefaultHTML) > 0 {
		widthStr := strconv.Itoa(config.Size[0])
		heightStr := strconv.Itoa(config.Size[1])
		file, err := os.Create(rootPath + "/output/" + widthStr + "x" + heightStr + "/index.html")
		if err != nil {
			return err
		}
		for _, s := range config.DefaultHTML {
			file.WriteString(s)
		}
	}

	return nil
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

// func writeCSSFromConfigurationFile(cssFilePath string, configPath string) {

// 	file, err := os.OpenFile(outpath, os.O_RDWR, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err := file.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	w := bufio.NewWriter(file)
// 	var i int
// 	for i < len(fileCSSTokens) {
// 		if _, err := w.Write([]byte(fileCSSTokens[i].Value)); err != nil {
// 			panic(err)
// 		}
// 		i++
// 	}

// 	if err = w.Flush(); err != nil {
// 		panic(err)
// 	}
// }
