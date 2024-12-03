// A collection of functions for navigating, interpreting, and validating folder structure
package validation

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

type BannerFolder struct {
	Name       string
	Subfolders []string
}

func ProjectValid(projectPath string) bool {
	/*
		Tests the validty of a project folder structure

		Note: An empty input folder IS allowed
		Note: Folders that are not named with [size]x[size] will be ignored, citing build flexibility

		This function will check for a valid project structure in the following format:

		root
			input/
			output/
				...at least one valid banner folder

	*/
	var inputExists bool = false
	var outputExists bool = false

	//read project path for 'input' and 'output' folders
	filepath.WalkDir(projectPath, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		//very opinionated, may want to add configuration here?
		if d.Name() == "input" {
			inputExists = true
		}
		if d.Name() == "output" {
			outputExists = true
		}

		if inputExists && outputExists {
			return fmt.Errorf("input and output found, exiting early")
		}

		return nil
	})

	if inputExists && outputExists {
		return FolderValid(projectPath)
	} else {
		return false
	}
}

func FolderValid(path string) bool {
	/*
		Tests the validity of a banner folder structure

		The function returns true if the following holds:

		Any subfolder of a banner parent folder named in the format [integer]x[integer] contains:
			-src/main.js
			-styles/main.css
			index.html

		If this is untrue return false

		Note: this function will ignore other folders intentionally, allowing for build flexibility
	*/
	var bannerFolders map[string]BannerFolder = getBannerFolders(path)

	for _, v := range bannerFolders {
		if !checkSubfolderStructure(v) {
			return false
		}
	}

	return true
}

func getBannerFolders(path string) map[string]BannerFolder {
	//Accepts a path string and returns a map of all banner sub-folders, along with array lists of their subfolders in the format:
	//Folder Name: digits x digits
	//	list of subfolders: [path,path]

	var bannerFolders map[string]BannerFolder = make(map[string]BannerFolder)
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dirNameMatch, dirNameMatchErr := regexp.MatchString(`\d+x\d+`, d.Name())
		bytePath := []byte(p)
		pathMatchObj, pathMatchErr := regexp.Compile(`\d+x\d+`)

		//handle errors
		if dirNameMatchErr != nil {
			fmt.Println(dirNameMatchErr)
			return dirNameMatchErr
		}

		if pathMatchErr != nil {
			fmt.Println(pathMatchErr)
			return pathMatchErr
		}

		pathLoc := pathMatchObj.FindIndex(bytePath)

		if d.IsDir() && dirNameMatch {
			//if not exists, set new empty bannerfolder in map
			if _, ok := bannerFolders[d.Name()]; !ok {
				var bf BannerFolder
				bf.Name = d.Name()
				bannerFolders[d.Name()] = bf
			}
		} else if pathLoc != nil {
			pathArr := strings.Split(p, "")
			foundBfolderName := strings.Join(pathArr[pathLoc[0]:pathLoc[1]], "")
			if entry, ok := bannerFolders[foundBfolderName]; ok {
				entry.Subfolders = append(entry.Subfolders, p)
				bannerFolders[foundBfolderName] = entry
			}
		}

		return nil
	})

	return bannerFolders
}

func checkSubfolderStructure(folder BannerFolder) bool {
	//checks the subfolder structure of a passed in BannerFolder struct
	var needs = map[string]bool{
		"src/main.js":     false,
		"styles/main.css": false,
		"index.html":      false,
	}

	//Note: this is O(n^2), can probably find a more efficient solution w/ hashmap
	for key := range needs {
		for _, e := range folder.Subfolders {
			if strings.Contains(e, key) {
				needs[key] = true
				break
			}
		}
	}

	for _, v := range needs {
		if !v {
			return false
		}
	}

	return true
}
