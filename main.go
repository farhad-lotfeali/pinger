package main

import (
	"fmt"
	"os"

	"github.com/farhad-lotfeali/pinger/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "app [command]",
		Short: "run pinger in back ground",
		Run: func(cmd *cobra.Command, arg []string) {
			fmt.Printf("please use %s [pinger|server]", os.Args[0])
		},
	}

	rootCmd.AddCommand(cmd.NewPinger())
	rootCmd.AddCommand(cmd.NewServer())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v", err)
	}
}
