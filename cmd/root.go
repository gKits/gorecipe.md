package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gkits/gorecipe.md/internal/recipe"
	"github.com/spf13/cobra"
)

var (
	outPath         string
	customTemplate  string
	withHugoHeaders bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorecipe.md",
	Short: "Scrape your favorie recipes and convert them to markdown.",
	Long: `gorecipe.md is small cli tool that scrapes recipes from the web and converts them to markdown.
It also supports the usage of Hugo by optionally adding the required headers for certain meta data.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		if _, err := url.Parse(args[0]); err != nil {
			return fmt.Errorf("positional argument has to be a valid url: %w", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		w := os.Stdout
		if outPath != "" {
			f, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("failed to open file at '%s': %e", outPath, err)
				os.Exit(1)
			}
			defer f.Close()
			w = f
		}

		if err := recipe.Convert(w, args[0], recipe.WithHugoHeaders(withHugoHeaders)); err != nil {
			fmt.Println()
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "", "Write the markdown result to a file.")
	rootCmd.Flags().StringVar(&customTemplate, "tmpl", "", "Use a custom template for the markdown file.")
	rootCmd.Flags().BoolVar(&withHugoHeaders, "hugo", false, "Add headers to the markdown file for usage with Hugo.")
}
