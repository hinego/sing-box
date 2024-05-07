package main

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/sagernet/sing-box/log"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/json"
	"github.com/sagernet/sing/common/rw"

	"github.com/spf13/cobra"
)

var commandMerge = &cobra.Command{
	Use:   "merge <output>",
	Short: "Merge configurations",
	Run: func(cmd *cobra.Command, args []string) {
		err := merge(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	mainCommand.AddCommand(commandMerge)
}

func merge(outputPath string) error {
	mergedOptions, err := readConfigAndMerge()
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(mergedOptions)
	if err != nil {
		return E.Cause(err, "encode config")
	}
	if existsContent, err := os.ReadFile(outputPath); err != nil {
		if string(existsContent) == buffer.String() {
			return nil
		}
	}
	err = rw.WriteFile(outputPath, buffer.Bytes())
	if err != nil {
		return err
	}
	outputPath, _ = filepath.Abs(outputPath)
	os.Stderr.WriteString(outputPath + "\n")
	return nil
}
