package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gkits/gorecipe.md/internal/recipe"
	"github.com/gkits/gorecipe.md/pkg/version"
	"github.com/spf13/cobra"
)

var (
	outPath         string
	customTemplate  string
	forced          bool
	withHugoHeaders bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorecipe.md",
	Short: "Scrape your favorie recipes and convert them to markdown.",
	Long: `gorecipe.md is small cli tool that scrapes recipes from the web and converts them to markdown.
It also supports the usage of Hugo by optionally adding the required headers for certain meta data.`,
	Version: version.Version(),
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
				fmt.Printf("failed to open file at '%s': %s\n", outPath, err)
				os.Exit(1)
			}
			// FIX: defer will not run due to os.Exit(1)
			defer f.Close()
			w = f
		}

		if err := recipe.MDScrape(
			w,
			args[0],
			recipe.WithTemplate(customTemplate),
			recipe.WithForced(forced),
			recipe.WithHugoHeaders(withHugoHeaders),
		); err != nil {
			if outPath != "" {
				if err := os.Remove(outPath); err != nil {
					fmt.Printf("failed to delete empty file: %s\n", err)
				}
			}
			fmt.Printf("failed to scrape recipe: %s\n", err)
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
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "", "path to output file")
	rootCmd.Flags().StringVar(&customTemplate, "tmpl", "", "custom markdown template")
	rootCmd.Flags().BoolVarP(&forced, "force", "f", false, "force markdown by ignoring missing recipe parts")
	rootCmd.Flags().BoolVar(&withHugoHeaders, "hugo", false, "add hugo headers")
}
