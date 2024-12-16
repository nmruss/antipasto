package filemgmt

import (
	"bufio"
	"os"

	"github.com/gorilla/css/scanner"
)

// NOTE: at some point maybe make this safer with our own tokenType?
type CSSToken struct {
	Type  string
	Value string
}

// stores property inserts / updates from image data
// NOTE: maybe these should be more restricted?
type CSSPropertyInsert struct {
	Type         string
	ParentName   string
	PropertyName string
	Value        string
}

// splits by curlybrace, returning the contained curlybrace
func splitByCurlyBrace(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i] == '}' {
			return i + 1, data[:i+1], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

// Takes in a CSS file path, opens the file and
// returns its contents as an array of CSSTokens
func TokenizeCSSFromFile(filepath *string) []CSSToken {
	var cssTokens []CSSToken

	file, err := os.OpenFile(*filepath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewScanner(file)
	r.Split(splitByCurlyBrace)
	for r.Scan() {
		s := scanner.New(r.Text())
		tok := s.Next()
		for tok.Type != scanner.TokenEOF && tok.Type != scanner.TokenError {
			//NOTE: if needed to deal with replacement UTF-8 character
			// if strings.ContainsRune(tok.Value, '\uFFFD') {
			// 	break
			// }
			var cssTok CSSToken
			cssTok.Type = tok.Type.String()
			cssTok.Value = tok.Value
			cssTokens = append(cssTokens, cssTok)
			tok = s.Next()
		}
	}
	return cssTokens
}

// Updates a []CSSToken list based on current []CSSToken list + []CSSPropertyInsert
func UpdateCSSTokenList(filepath *string, currentProperties *[]CSSToken, updates *[]CSSPropertyInsert) {
	currProps := *currentProperties
	u := *updates

	//this is O(n^2), could be better possibly with a tree or using a map for constant time lookup
	for _, update := range u {
		right := 0
		propExists := false
		selectorExists := false

		for right < len(currProps) {
			//if there is an existing selector (will be of type HASH if an id)
			if currProps[right].Value == update.ParentName {
				selectorExists = true

				//search for the next IDENT token with the correct property name
				for right < len(currProps) && currProps[right].Value != "}" {
					if currProps[right].Type == "IDENT" && currProps[right].Value == update.PropertyName {
						propExists = true
						for right < len(currProps) && currProps[right].Value != "}" {
							///search for the next token of update type, update it
							if currProps[right].Type == update.Type {
								currProps[right].Value = update.Value
								break
							}
							right++
						}
					}
					right++
				}

				//If you reach a "}" before you see a matching property name, insert
				//an IDENT, CHAR ':' and DIMENSION token into the currProps list
				//along with a char ';' to complete the valid CSS insert
				if !propExists {
					newTab := CSSToken{Type: "S", Value: "      "}
					newIdentifier := CSSToken{Type: "IDENT", Value: update.PropertyName}
					newColon := CSSToken{Type: "CHAR", Value: ":"}
					newDimension := CSSToken{Type: update.Type, Value: update.Value}
					newSemicolon := CSSToken{Type: "CHAR", Value: ";"}
					newReturn := CSSToken{Type: "S", Value: "\n"}
					currProps = append(currProps[:right], append([]CSSToken{newTab, newIdentifier, newColon, newDimension, newSemicolon, newReturn}, currProps[right:]...)...)
				}
			}
			right++
		}

		//if you've reached a "}" and the selector does not exist
		if !selectorExists {
			var selectorType string

			switch update.ParentName[0] {
			case '#':
				selectorType = "HASH"
			case '.':
				selectorType = "DOT"
			}

			//newIdentifier := CSSToken{Type: selectorType, Value: update.PropertyName}
			currProps = append(currProps, []CSSToken{
				{Type: selectorType, Value: update.ParentName},
				{Type: "CHAR", Value: "{"},
				{Type: "S", Value: "\n"},
				{Type: "S", Value: "      "},
				{Type: "IDENT", Value: update.PropertyName},
				{Type: "CHAR", Value: ":"},
				{Type: update.Type, Value: update.Value},
				{Type: "CHAR", Value: ";"},
				{Type: "S", Value: "\n"},
				{Type: "CHAR", Value: "}"},
			}...)
		}
	}

	*currentProperties = currProps
}

func WriteCSS(filepath *string, updates *[]CSSPropertyInsert, outpath string) {
	// takes an existing css file path
	// and a list of CSS token updates
	// wries the listed CSS token updates to the file at path
	file, err := os.OpenFile(outpath, os.O_RDWR, 0644)
	fileCSSTokens := TokenizeCSSFromFile(filepath)
	UpdateCSSTokenList(filepath, &fileCSSTokens, updates)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(file)
	var i int
	for i < len(fileCSSTokens) {
		if _, err := w.Write([]byte(fileCSSTokens[i].Value)); err != nil {
			panic(err)
		}
		i++
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

// func printBuffer(buf *[]byte, asString bool) {
// 	//Utility function that prints a byte buffer as bytes or a string
// 	//followed by a new line character
// 	buffer := *buf
// 	for n := 0; n < len(buffer); n++ {
// 		if asString {
// 			fmt.Print(string(buffer[n]))
// 		} else {
// 			fmt.Print(buffer[n])
// 		}
// 		if buffer[n] == 0 {
// 			fmt.Println()
// 			break
// 		}
// 	}
// }
