package configuration

import (
	"bufio"
	"os"
)

// Contains configuration items, such as the default write config when generating a new project
type APConfig struct {
	DefaultHTML    string
	DefaultJS      string
	DefaultStyling string
}

type scanState int

const (
	html scanState = iota
	js
	css
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
		}

		switch state {
		case html:
			returnConfig.DefaultHTML += currentLine
		case css:
			returnConfig.DefaultStyling += currentLine
		case js:
			returnConfig.DefaultJS += currentLine
		}
	}

	return returnConfig
}
