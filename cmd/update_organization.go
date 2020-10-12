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

// createOrganizationCmd represents the createOrganization command
var updateOrganizationCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an organization",
	Long: `Example:

packet organization update --id [organization_UUID] --name [new_name]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &packngo.OrganizationUpdateRequest{}

		if name != "" {
			req.Name = &name
		}

		if description != "" {
			req.Description = &description
		}

		if twitter != "" {
			req.Twitter = &twitter
		}

		if logo != "" {
			req.Logo = &logo
		}

		org, _, err := apiClient.Organizations.Update(organizationID, req)
		if err != nil {
			return errors.Wrap(err, "Could not update Organization")
		}

		data := make([][]string, 1)

		data[0] = []string{org.ID, org.Name, org.Created}
		header := []string{"ID", "Name", "Created"}

		return output(org, header, &data)
	},
}

func init() {
	updateOrganizationCmd.Flags().StringVarP(&organizationID, "id", "i", "", "Organization ID")
	updateOrganizationCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the organization")
	updateOrganizationCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the organization")
	updateOrganizationCmd.Flags().StringVarP(&website, "website", "w", "", "Website of the organization")
	updateOrganizationCmd.Flags().StringVarP(&twitter, "twitter", "t", "", "Twitter URL of the organization")
	updateOrganizationCmd.Flags().StringVarP(&logo, "logo", "l", "", "Logo URL of the organization")

	_ = updateOrganizationCmd.MarkFlagRequired("id")
}
