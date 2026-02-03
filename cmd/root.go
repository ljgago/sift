package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   = "(untracked)"
	commit    = "(untracked)"
	buildDate = "(untracked)"
	system    = "(untracked)"
	arch      = "(untracked)"
)

type Flags struct {
	Version bool
}

// rootCmd represents the root
var rootCmd = &cobra.Command{
	Use:   "sift",
	Short: "An npm search packages viewer",
	Long:  `An npm search packages viewer`,
	Run: func(cmd *cobra.Command, args []string) {
		rootRun(cmd, args)
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "version of api-viewer")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseFlags(cmd *cobra.Command) (*Flags, error) {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		return nil, fmt.Errorf("Error getting version flag -> %w", err)
	}

	return &Flags{
		Version: version,
	}, nil
}

func rootRun(cmd *cobra.Command, _ []string) {
	flags, err := parseFlags(cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if flags.Version {
		fmt.Printf("api-viewer version %s\nCommit: %s\nBuild Date: %s\nOS/Arch: %s/%s\n", version, commit, buildDate, system, arch)
		return
	}

	cmd.Help()
}
