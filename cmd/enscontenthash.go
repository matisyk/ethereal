// Copyright © 2017-2019 Weald Technology Trading
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

// ensContenthashCmd represents the ens contenthash command
var ensContenthashCmd = &cobra.Command{
	Use:   "contenthash",
	Short: "Manage ENS contenthash entries",
	Long:  `Set and obtain Ethereum Name Service contenthash information`,
}

func init() {
	ensCmd.AddCommand(ensContenthashCmd)
}

func ensContenthashFlags(cmd *cobra.Command) {
	ensFlags(cmd)
}
