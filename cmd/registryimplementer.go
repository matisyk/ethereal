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
	"github.com/spf13/cobra"
)

var registryImplementerInterface string
var registryImplementerAddressStr string

// registryImplementerCmd represents the registry implementer command
var registryImplementerCmd = &cobra.Command{
	Use:   "implementer",
	Short: "Manage ERC-1820 registry implementers",
	Long:  `Set and obtain ERC-1820 registry implementer information`,
}

func init() {
	registryCmd.AddCommand(registryImplementerCmd)
}

func registryImplementerFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&registryImplementerInterface, "interface", "", "interface against which to operate (e.g. ERC777TokensRecipient)")
	cmd.Flags().StringVar(&registryImplementerAddressStr, "address", "", "address against which to operate (e.g. wealdtech.eth)")
}
