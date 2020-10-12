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
	"fmt"
	"os"
	"runtime"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// apiClient client
	apiClient   packngo.Client
	cfgFile     string
	isJSON      bool
	isYaml      bool
	packetToken string

	includes *[]string // nolint:unused
	excludes *[]string // nolint:unused

)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "packet",
	Short:             "Command line interface for Equinix Metal",
	Long:              `Command line interface for Equinix Metal`,
	DisableAutoGenTag: true,
	PersistentPreRunE: apiConnect,
}

func apiConnect(cmd *cobra.Command, args []string) error {
	if packetToken == "" {
		return fmt.Errorf("Equinix Metal authentication token not provided. Please either set the 'PACKET_TOKEN' environment variable or create a JSON or YAML configuration file.")
	}
	client, err := packngo.NewClientWithBaseURL(consumerToken, packetToken, nil, apiURL)
	if err != nil {
		return errors.Wrap(err, "Could not create Client")
	}
	client.UserAgent = fmt.Sprintf("packet-cli/%s %s", Version, client.UserAgent)
	apiClient = *client
	return nil
}

func apiToken() string {
	return os.Getenv(apiTokenEnvVar)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to JSON or YAML configuration file")

	rootCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	rootCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")

	includes = rootCmd.PersistentFlags().StringSlice("include", nil, "Comma seperated Href references to expand in results, may be dotted three levels deep")
	excludes = rootCmd.PersistentFlags().StringSlice("exclude", nil, "Comma seperated Href references to collapse in results, may be dotted three levels deep")

	rootCmd.Version = Version
}

// listOptions creates a ListOptions using the includes and excludes persistent
// flags. When not defined, the defaults given will be supplied.
func listOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions {
	listOptions := &packngo.ListOptions{
		Includes: defaultIncludes,
		Excludes: defaultExcludes,
	}
	if rootCmd.Flags().Changed("include") {
		listOptions.Includes = *includes
	}
	if rootCmd.Flags().Changed("exclude") {
		listOptions.Excludes = *excludes
	}

	return listOptions
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".packet-cli")
		viper.AddConfigPath(userHomeDir())
	}

	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		panic(fmt.Errorf("Could not read config: %s", err))
	}

	if viper.GetString("token") != "" {
		packetToken = viper.GetString("token")
	} else {
		packetToken = apiToken()
	}
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
