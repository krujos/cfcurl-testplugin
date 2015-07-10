package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/cfcurl"
)

//Run a command
func (cmd *TestCmd) Run(cliConnection plugin.CliConnection, args []string) {
	out, _ := cfcurl.Curl(cliConnection, "/v2/apps")
	fmt.Println(out)

	out, _ = cfcurl.CurlDepricated(cliConnection, "/v2/domains")
	fmt.Println(out)
}

//TestCmd does a test
type TestCmd struct {
}

//GetMetadata returns metatada
func (cmd *TestCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "test",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "testplugin",
				HelpText: "testplugin",
				UsageDetails: plugin.Usage{
					Usage: "cf testplugin",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(TestCmd))
}
