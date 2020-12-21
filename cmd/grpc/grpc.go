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
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/mihmatache/hello-world/pkg/certificates"
	grpcHello "github.com/mihmatache/hello-world/pkg/grpc"
)

var (
	authority string
	rootCA    string
	cert      string
	key       string
	isTls     bool
	isMtls    bool
	timeout   int
)

func init() {
	Cmd.AddCommand(server)
	Cmd.AddCommand(client)
	client.Flags().StringVar(&authority, "authority", "", "authority to use when calling")
	Cmd.PersistentFlags().StringVar(&rootCA, "rootCA", "", "authority to use when calling")
	Cmd.PersistentFlags().StringVar(&cert, "cert", "", "authority to use when calling")
	Cmd.PersistentFlags().StringVar(&key, "key", "", "authority to use when calling")
	Cmd.PersistentFlags().BoolVar(&isTls, "tls", false, "Set security to TLS")
	Cmd.PersistentFlags().BoolVar(&isMtls, "mtls", false, "Set security to mTLS")
	client.Flags().IntVar(&timeout, "timeout", 30, "call timeout")
}

// Cmd represents the grpc command
var Cmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC type connections",
	Long:  `gRPC type connections will be used. You can either call it as a `,
}

var server = &cobra.Command{
	Use:   "server",
	Short: "gRPC server start",
	Long:  `Start a gRPC hello server`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		options := make([]grpc.ServerOption, 0)
		if isTls || isMtls {
			conf := &tls.Config{}
			err := certificates.SetTLSCertificates(conf, cert, key)
			if err != nil {
				panic(err)
			}
			if isTls {
				conf.ClientAuth = tls.NoClientCert
			} else {
				conf.ClientAuth = tls.RequireAndVerifyClientCert
			}
			options = append(options, grpc.Creds(certificates.MakeCredentials(conf)))
		}
		err = grpcHello.StartServer(port, options)
		if err != nil {
			panic(err)
		}
	},
}

var client = &cobra.Command{
	Use:   "client",
	Short: "gRPC client start",
	Long:  `Start a gRPC hello client`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			panic(fmt.Errorf("Please provide a single address to the command"))
		}

		client := grpcHello.NewClient(args[0])
		if authority != "" {
			client.WithOption(grpc.WithAuthority(authority))
		}
		if !isTls && !isMtls {
			client.WithOption(grpc.WithInsecure())
		} else {
			rootCert, err := ioutil.ReadFile(rootCA)
			if err != nil {
				panic(err)
			}
			conf := &tls.Config{}
			err = certificates.SetRootTLSCert(conf, rootCert)
			if err != nil {
				panic(err)
			}
			if isMtls {
				certificates.SetTLSCertificates(conf, cert, key)
			}
			client.WithOption(grpc.WithTransportCredentials(certificates.MakeCredentials(conf)))
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		b, err := grpcHello.ClientCall(ctx, client)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	},
}
