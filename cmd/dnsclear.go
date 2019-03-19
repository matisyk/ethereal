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
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

// dnsClearCmd represents the dns clear command
var dnsClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all records for a DNS zone",
	Long: `Clear all records for a DNS zone.  For example to cler all records for wealdtech.eth:

    ethereal dns clear --domain=wealdtech.eth --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(dnsDomain != "", quiet, "--domain is required")
		if !strings.HasSuffix(dnsDomain, ".") {
			dnsDomain = dnsDomain + "."
		}
		dnsDomain = ens.NormaliseDomain(dnsDomain)
		outputIf(verbose, fmt.Sprintf("DNS domain is %s", dnsDomain))
		ensDomain := strings.TrimSuffix(dnsDomain, ".")
		outputIf(verbose, fmt.Sprintf("ENS domain is %s", ensDomain))
		domainHash := ens.NameHash(ensDomain)
		outputIf(verbose, fmt.Sprintf("ENS domain hash is 0x%x", domainHash))

		// Obtain the registry contract
		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Obtain owner for the domain
		domainOwner, err := registryContract.Owner(nil, domainHash)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")

		cli.Assert(bytes.Compare(domainOwner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, "Owner is not set")
		outputIf(verbose, fmt.Sprintf("Domain owner is %s", ens.Format(client, &domainOwner)))

		// Obtain resolver for the domain
		resolverAddress, err := ens.Resolver(registryContract, ensDomain)
		cli.ErrCheck(err, quiet, fmt.Sprintf("No resolver registered for %s", dnsDomain))
		resolverContract, err := ens.DNSResolverContractByAddress(client, resolverAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", dnsDomain))
		outputIf(debug, fmt.Sprintf("Resolver contract is at %s", ens.Format(client, &resolverAddress)))

		// Build the transaction
		opts, err := generateTxOpts(domainOwner)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := resolverContract.ClearDNSZone(opts, domainHash)
		cli.ErrCheck(err, quiet, "Failed to create transaction")
		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
			os.Exit(_exit_success)
		}

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":   "dns",
			"command": "clear",
			"domain":  dnsDomain,
		})
	},
}

func init() {
	dnsCmd.AddCommand(dnsClearCmd)
	dnsFlags(dnsClearCmd)
	addTransactionFlags(dnsClearCmd, "the owner of the domain")
}
