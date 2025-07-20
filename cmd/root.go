package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"stashman/core"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stashman",
	Short: "A CLI tool for managing code snippets from annotated source files.",
	Long: `Stashman is a terminal-based code snippet manager that parses source files for specially 
marked comments (@stash ... @endstash) and stores the enclosed code blocks in a local JSON database.

Commands:

  stashman scan [dir|--file <path>]  Scans a directory (default) or a single file for @stash comments
  stashman list                      Lists all saved snippets with IDs, titles, tags, and language
  stashman get <id>                  Prints the full content of a snippet by ID
  stashman search [flags]            Searches snippets by tag, title, or language
  stashman update                    Refreshes snippets that have changed in their original files

Global Flags:

  --stash-file <path>                Path to the stash JSON file (default: .stash.json)
  --ext <extension>                  File extension to scan (e.g., --ext=go)
  --lines <n>                        Max lines to capture if @endstash is not found
  --no-gitignore                     By default, Stashman respects .gitignore and skips ignored files and directories. To disable this behavior, pass this flag.

Snippets must be annotated in source files using:

  // @stash title=Example Snippet tag=go,db
     ...your code here...
  // @endstash

Snippets are stored with metadata such as title, tags (multiple tags separated by commas), language, and source location.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stashman.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
    	core.LoadOrInitProject()
    	return nil
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


