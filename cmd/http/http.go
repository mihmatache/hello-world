/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package http

import (
	"github.com/spf13/cobra"

	httpHello "github.com/mihmatache/hello-world/pkg/http"
)

var (

)

func init() {
	Cmd.AddCommand(server)
}

// Cmd represents the grpc command
var Cmd = &cobra.Command{
	Use:   "http",
	Short: "http type connections",
	Long:  `http type connections will be used.`,
}

var server = &cobra.Command{
	Use:   "server",
	Short: "HTTP server start",
	Long:  `Start a HTTP hello server`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")

		err = httpHello.StartServer(port)
		if err != nil {
			panic(err)
		}
	},
}
