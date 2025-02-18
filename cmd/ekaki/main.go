package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/aethiopicuschan/ekaki/pkg/ekaki"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "ekaki [source] [output]",
	Long:         "Ekaki is a simple CLI tool to convert images.",
	Args:         cobra.MatchAll(cobra.MinimumNArgs(2), cobra.MaximumNArgs(2)),
	RunE:         run,
	SilenceUsage: true,
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		rootCmd.Version = bi.Main.Version
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	sp := args[0]
	op := args[1]

	source, err := os.Open(sp)
	if err != nil {
		return
	}
	defer source.Close()

	expr := filepath.Ext(op)[1:]
	t, err := ekaki.TargetFromExpr(expr)
	if err != nil {
		return
	}

	output, err := ekaki.Convert(source, t)
	if err != nil {
		return
	}

	file, err := os.Create(op)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(output))

	return
}

func main() {
	rootCmd.Execute()
}
