// A collection of functions for navigating, interpreting, and validating folder structure
package navigation

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

func FolderValid(path string) bool {
	/*
		Tests the validity of a banner folder structure

		The function returns true if the following holds:

		Any subfolder of a banner parent folder named in the format 'integer'x'integer' contains:
			-src/main.js
			-styles/main.css
			index.html

		If this is untrue return false

		Note: this function can ignore other folders intentionally, allowing for build flexibility
	*/
	var bannerFolders map[string]BannerFolder = getBannerFolders(path)

	for _, v := range bannerFolders {
		if !checkSubfolderStructure(v) {
			return false
		}
	}

	return true
}

func checkSubfolderStructure(folder BannerFolder) bool {
	//checks the subfolder structure of a passed in BannerFolder struct
	var needs = map[string]bool{
		"src/main.js":     false,
		"styles/main.css": false,
		"index.html":      false,
	}

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
