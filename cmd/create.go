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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ericogr/cdcf/docker"
)

var EncodeBase64 bool
var Address string
const DEFAULT_ADDRESS = "https://index.docker.io/v1/"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a docker config file",
	Long: `Using credencials to output a docker config file.`,
	Args: cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		var address string

		if len(args) == 2 {
			address = DEFAULT_ADDRESS
		} else {
			address = args[2]
		}

		if EncodeBase64 {
			fmt.Println(docker.CreateDockerConfigEncoded(args[0], args[1], address))
		} else {
			fmt.Println(docker.CreateDockerConfig(args[0], args[1], address))
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.PersistentFlags().BoolVarP(&EncodeBase64, "encodeBase64", "e", true, "encode output in base64")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", DEFAULT_ADDRESS, "docker registry url")
	createCmd.Flags().BoolVarP(&EncodeBase64, "encodeBase64", "e", true, "encode output in base64")
}
