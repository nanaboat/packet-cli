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
	"github.com/packethost/packngo"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var (
	description string
	locked      bool
)

// updateDeviceCmd represents the updateDevice command
var updateDeviceCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a device",
	Long: `Example:

packet device update --id [device_UUID] --hostname [new_hostname]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &packngo.DeviceUpdateRequest{}

		if hostname != "" {
			req.Hostname = &hostname
		}

		if description != "" {
			req.Description = &description
		}

		if userdata != "" {
			req.UserData = &userdata
		}

		if locked {
			req.Locked = &locked
		}

		if len(tags) > 0 {
			req.Tags = &tags
		}

		if alwaysPXE {
			req.AlwaysPXE = &alwaysPXE
		}

		if ipxescripturl != "" {
			req.IPXEScriptURL = &ipxescripturl
		}

		if customdata != "" {
			req.CustomData = &customdata
		}

		device, _, err := apiClient.Devices.Update(deviceID, req)
		if err != nil {
			return errors.Wrap(err, "Could not update Device")
		}

		header := []string{"ID", "Hostname", "OS", "State"}
		data := make([][]string, 1)
		data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State}

		return output(device, header, &data)
	},
}

func init() {
	updateDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "UUID of the device")
	updateDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "Hostname")
	updateDeviceCmd.Flags().StringVarP(&description, "description", "d", "", "Description for the device")
	updateDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "User data")
	updateDeviceCmd.Flags().BoolVarP(&locked, "locked", "l", false, "Lock device")
	updateDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Tags for the device --tags="tag1,tag2"`)
	updateDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, `--alaways-pxe or -a`)
	updateDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "s", "", "URL to the iPXE script")
	updateDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "Custom data")

	_ = updateDeviceCmd.MarkFlagRequired("id")

}
