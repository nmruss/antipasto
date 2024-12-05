package cmd

import (
	"fmt"
	"nmruss/antipasto/validation"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Determines project validity based on folder structure",
	Long: `Validate determines validity of a project based on its folder structure
A valid project must have the following structure:
/
 input/
 output/
  styles/
   main.css
  src/
   main.js
  index.html

Note: So long as these files exist, you have a valid project.
For flexibility and ease of use, you can add whatever other 
files are needed into your output folder.
Empty input folders are also allowed.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a project path to validate")
		} else {
			fmt.Println(validation.IsProjectValid(args[0]))
		}
	},
}
