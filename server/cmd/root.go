package main

import (
	"fmt"
	"os"

	"github.com/fatedier/frp/pkg/auth"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/util"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/fatedier/frp/server"

	"github.com/spf13/cobra"
)

const (
	CfgFileTypeIni = iota
	CfgFileTypeCmd
)

var (
 cfgFile string 	
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "conf","c", "./application.toml","config file" )
}

var rootCmd = &cobra.Command{
	Use:   "trains",
	Short: "trains is the server of tranin (https://github.com/fatedier/frp)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

