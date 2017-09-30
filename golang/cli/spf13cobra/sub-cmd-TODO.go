package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "mycli"}

func main() {

	CmdCustomer.AddCommand(CmdCustomerGet)
	CmdCustomer.AddCommand(CmdCustomerPut)
	RootCmd.AddCommand(CmdCustomer)

	CmdApp.AddCommand(CmdAppGet)
	CmdApp.AddCommand(CmdAppPut)
	RootCmd.AddCommand(CmdApp)

	RootCmd.Execute()
}

// command with mandatory sub-commands
var CmdCustomer = &cobra.Command{
	Use:   "customer",
	Short: "short message for customer",
	Long:  "long message for customer",
}

var CmdCustomerGet = &cobra.Command{
	Use:   "get",
	Short: "short message for get customer",
	Long:  "long message for get customer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get customer")
	},
}

var CmdCustomerPut = &cobra.Command{
	Use:   "put",
	Short: "short message for put customer",
	Long:  "long message for put customer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing put customer")
	},
}

//
// command with optional sub-commands
var CmdApp = &cobra.Command{
	Use:   "app",
	Short: "short message for app",
	Long:  "long message for app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing app")
	},
}

// sub command having 1 mandatory params
var CmdAppGet = &cobra.Command{
	Use:   "get",
	Short: "short message for get app",
	Long:  "long message for get app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get app")
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing arguments")
		}
		return nil
	},
}

// sub command having 2 mandatory params
var CmdAppPut = &cobra.Command{
	Use:   "put",
	Short: "short message for put app",
	Long:  "long message for put app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing put app")
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("missing arguments")
		}
		return nil
	},
}
