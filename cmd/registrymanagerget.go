// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
	erc1820 "github.com/wealdtech/go-erc1820"
)

// registryManagerGetCmd represents the registry manager get command
var registryManagerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the address of an ERC-1820 interface manager",
	Long: `Obtain the address of an manager registered with the ERC-1820 registry for a given address.  For example:

    ethereal registry manager get --address=0x1234...5678

Note that this will always return the managing address if possible.  This means that if a manager is not set it will return the provided address

In quiet mode this will return 0 if the manager was obtained without error, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		address, err := ens.Resolve(client, registryManagerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		registry, err := erc1820.NewRegistry(client)
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		manager, err := registry.Manager(&address)
		cli.ErrCheck(err, quiet, "failed to obtain manager")

		if *manager == ens.UnknownAddress {
			manager = &address
		}
		if !quiet {
			name, _ := ens.ReverseResolve(client, manager)
			if name == "" {
				name = manager.Hex()
			}
			fmt.Printf("%s\n", name)
		}
		os.Exit(0)
	},
}

func init() {
	initAliases(registryManagerGetCmd)
	registryManagerFlags(registryManagerGetCmd)
	registryManagerCmd.AddCommand(registryManagerGetCmd)
}
