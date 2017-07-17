package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	CmdGet.AddCommand(CmdGetContact)
	RootCmd.AddCommand(CmdGet)
	RootCmd.AddCommand(CmdPut)
	RootCmd.Execute()
}

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "short message for get",
	Long:  "long message for get",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get")
	},
}

var CmdGetContact = &cobra.Command{
	Use:   "contact",
	Short: "short message for get contact",
	Long:  "long message for get contact",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get contact")
	},
}

var CmdPut = &cobra.Command{
	Use:   "put",
	Short: "short message for put",
	Long:  "long message for put",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing put")
	},
}

var RootCmd = &cobra.Command{Use: "mycli"}
