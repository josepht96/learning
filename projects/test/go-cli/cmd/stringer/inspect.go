package stringer

import (
	"fmt"

	"github.com/josepht96/learning/projects/go-cli/pkg/stringer"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := stringer.Inspect(i, false)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
