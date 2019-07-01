package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rgreinho/gollaborators/gollaborators"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gollaborators",
	Short: "Generate your collaborator page.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check the number of arguments.
		if len(args) != 1 {
			log.Println("Requires a project owner/name.")
			os.Exit(0)
		}

		// Ensure a project owner/name was provided.
		project := strings.Split(args[0], "/")
		if len(project) != 2 {
			log.Println("Requires a project owner/name.")
			os.Exit(0)
		}

		// Get the maximum number of items per line.
		lineLength, err := cmd.Flags().GetInt("line-length")
		if err != nil {
			log.Println("Cannot read the \"line-length\" argument.")
		}

		// Retrieve the collaborators.
		if err := gollaborators.Retrieve(project[0], project[1], lineLength); err != nil {
			log.Printf("Error: %s", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("line-length", "l", 5, "Maximum number of collaborators per line.")
}
