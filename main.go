package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sydneyowl/shx8800-config-editor/cmd"
	"os"
)

func main() {
	cobra.MousetrapHelpText = ""
	if err := cmd.BaseCmd.Execute(); err != nil {
		fmt.Printf("程序无法启动: %v", err)
		os.Exit(-1)
	}
}
