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
package grpc

import (
	"fmt"

	"github.com/spf13/cobra"

	grpcHello "github.com/mihmatache/hello-world/pkg/grpc"
)

var (
	address string
	timeout int
)

func init(){
	Cmd.AddCommand(server)
	Cmd.AddCommand(client)
	client.Flags().StringVar(&address, "address", "127.0.0.1", "address of the gRPC server")
	client.Flags().IntVar(&timeout, "timeout", 30, "call timeout")
}

// Cmd represents the grpc command
var Cmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC type connections",
	Long: `gRPC type connections will be used. You can either call it as a `,
}


var server = &cobra.Command{
	Use:   "server",
	Short: "gRPC server start",
	Long: `Start a gRPC hello server`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		err = grpcHello.StartServer(port)
		if err != nil {
			panic(err)
		}
	},
}

var client = &cobra.Command{
	Use:   "client",
	Short: "gRPC client start",
	Long: `Start a gRPC hello client`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		b, err := grpcHello.ClientCall(address, port, timeout)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	},
}
