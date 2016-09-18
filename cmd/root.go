package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DebugMode bool

var RootCmd = &cobra.Command{
	Use:   "chronosctl",
	Short: "chronosctl",
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chronosctl v0.01")
	},
}

func InitRootCmd() {
	InitJobCmd()
	RootCmd.PersistentFlags().BoolVarP(&DebugMode, "debug", "d", false, "Will output all the outgoing http requests and responses")
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(JobCmd)
}
