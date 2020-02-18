package cmd

import(
	"fmt"
	"os"
	"path/filepath"
	"github.com/aferno/dupm/files"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "dupm is duplicate manager",
	Long: "Description: dupm is duplicate file finder and manager",
	Run: func(cmd *cobra.Command, args []string) {
		curDir, _ := filepath.Abs(".")
		fmt.Println("Starting dupm for", curDir)
		files.WalkInRoot(curDir)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	  	os.Exit(1)
	}
}
