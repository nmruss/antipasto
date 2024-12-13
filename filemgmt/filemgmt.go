package filemgmt

import (
	"bufio"
	"fmt"
	"io"
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
type CSSProperyInsert struct {
	ParentName   string
	PropertyName string
	Value        string
}

// Takes in a CSS file path and
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

	r := bufio.NewReader(file)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		s := scanner.New(string(buf))
		tok := s.Next()
		for tok.Type.String() != "EOF" {
			var cssTok CSSToken
			cssTok.Type = tok.Type.String()
			cssTok.Value = tok.Value
			cssTokens = append(cssTokens, cssTok)
			tok = s.Next()
		}
	}
	return cssTokens
}

func printBuffer(buf *[]byte, asString bool) {
	//Utility function that prints a byte buffer as bytes or a string
	//followed by a new line character
	buffer := *buf
	for n := 0; n < len(buffer); n++ {
		if asString {
			fmt.Print(string(buffer[n]))
		} else {
			fmt.Print(buffer[n])
		}
		if buffer[n] == 0 {
			fmt.Println()
			break
		}
	}
}

// Updates a []CSSToken list based on current []CSSToken list + []CSSPropertyInsert
func UpdateCSSTokenList(filepath *string, currentProperties *[]CSSToken, updates *map[string]CSSProperyInsert) {
	//NOTE: For now this will only update id's with top and left, width and height dimensions

	//step along the currentProperties list, if you find a token where Value == update.ParentName
	//start a pointer at the current index
	//step along the until you see a token that matches propertyName, or until you see a token that matches '}'
	//if you see a token that matches propertyName, continue to step along until you see a DIMENSION token
	//update this dimension token to equal update.Value, mark this value as updated in a map?
	currProps := *currentProperties
	u := *updates

	for i, currentProperty := range currProps {
		up, ok := u[currentProperty.Value]
		if ok {
			//left := i
			right := i
			for currProps[right].Value != "}" {
				if currProps[right].Type == "IDENT" && currProps[right].Value == up.PropertyName {
					//search for the next DIMENSION identfier and update the existing value
					for currProps[right].Type != "DIMENSION" {
						right++
					}

					if currProps[right].Type == "CHAR" && currProps[right].Value == "}" {
						//you reached a "}" before you saw a matching property name, insert
						//an IDENT, CHAR ':' and DIMENSION token into the currProps list
						break
					}

					if currProps[right].Type == "DIMENSION" && currProps[right].Value != up.Value {
						currProps[right].Value = up.Value
						break
					}
				}
				right++
			}
		}
	}

}

//func WriteCSS(){
//takes an array of type Property
// file, err := os.OpenFile(*filepath, os.O_RDWR, 0644)
// if err != nil {
// 	panic(err)
// }
// defer func() {
// 	if err := file.Close(); err != nil {
// 		panic(err)
// 	}
// }()

// r := bufio.NewReader(file)
// w := bufio.NewWriter(file)
// buf := make([]byte, 1024)
// for {
// 	n, err := r.Read(buf)
// 	if err != nil && err != io.EOF {
// 		panic(err)
// 	}

// 	if n == 0 {
// 		break
// 	}

// 	if _, err := w.Write(buf[:n]); err != nil {
// 		panic(err)
// 	}
// }

// if err = w.Flush(); err != nil {
// 	panic(err)
// }
//}
