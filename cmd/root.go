/*
Copyright © 2024 GABRIEL BRANDÃO contato@gabrielbrandao.net
*/
package cmd

import (
	"database/sql"
	"os"

	dbInfra "github.com/brandaogabriel7/full-cycle-arquitetura-hexagonal/adapters/db"
	"github.com/brandaogabriel7/full-cycle-arquitetura-hexagonal/application"
	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = dbInfra.NewProductDb(db)
var productService = application.NewProductService(productDb)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "full-cycle-arquitetura-hexagonal",
	Short: "Example CLI products application with Hexagonal Architecture",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.full-cycle-arquitetura-hexagonal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
