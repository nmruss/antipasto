package configuration

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Contains configuration items, such as the default write config when generating a new project
type APConfig struct {
	DefaultHTML    []string
	DefaultJS      []string
	DefaultStyling []string
	Size           []int
}

type scanState int

const (
	html scanState = iota
	js
	css
	size
	none
)

// parses an Antipasto configuration file, and returns an APConfig object
func ParseConfigurationFile(configPath string) APConfig {
	returnConfig := APConfig{}
	file, err := os.Open(configPath)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewScanner(file)
	var state scanState
	state = none
	r.Split(bufio.ScanLines)

	for r.Scan() {
		currentLine := r.Text()

		switch currentLine {
		case "AP_HTML_TEMPLATE":
			state = html
			continue
		case "AP_HTML_TEMPLATE_END":
			state = none
		case "AP_JS_TEMPLATE":
			state = js
			continue
		case "AP_JS_TEMPLATE_END":
			state = none
		case "AP_CSS_TEMPLATE":
			state = css
			continue
		case "AP_CSS_TEMPLATE_END":
			state = none
		case "AP_SIZE":
			state = size
			continue
		case "AP_SIZE_END":
			state = none
		}

		switch state {
		case html:
			returnConfig.DefaultHTML = append(returnConfig.DefaultHTML, currentLine+"\n")
		case css:
			returnConfig.DefaultStyling = append(returnConfig.DefaultStyling, currentLine+"\n")
		case js:
			returnConfig.DefaultJS = append(returnConfig.DefaultJS, currentLine+"\n")
		case size:
			sizeStrings := strings.Split(currentLine, "x")
			w, err := strconv.Atoi(sizeStrings[0])
			if err != nil {
				fmt.Println("error converting width string")
			}

			h, err := strconv.Atoi(sizeStrings[1])
			if err != nil {
				fmt.Println("error converting height string")
			}

			returnConfig.Size = append(returnConfig.Size, w)
			returnConfig.Size = append(returnConfig.Size, h)
		}
	}

	return returnConfig
}
