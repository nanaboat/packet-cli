// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var retrieveVpnCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves VPN service",
	Long: `Example:
	
Enable VPN service: 
packet vpn get --faciliy ewr1
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, _, err := apiClient.VPN.Get(facility, nil)
		if err != nil {
			return errors.Wrap(err, "Could not get VPN")
		}

		data := make([][]string, 1)

		data[0] = []string{config.Config}
		header := []string{"Config"}
		return output(config, header, &data)
	},
}

func init() {
	retrieveVpnCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility for which VPN config is to be retrieved")
	_ = retrieveVpnCmd.MarkFlagRequired("id")
}
