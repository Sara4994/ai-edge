/*
Copyright 2024. Open Data Hub Authors

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

package images

import (
	tea "github.com/charmbracelet/bubbletea"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/spf13/cobra"
)

var describeCmd = NewCmd(
	"describe <model-image-id> <model-version-name>",
	"Show details of an model image along with its params",
	`Show details of an model image along with its params
	This command allows you to check details about the model image which can be accessed by passing 
	model image id and model version name as arguments to the command
	`,
	cobra.ExactArgs(2),
	[]Flag{
		FlagNamespace.SetInherited(), FlagModelRegistryUrl.SetInherited(), FlagKubeconfig.SetInherited(),
		FlagParams,
	},
	SubCommandDescribe,
	func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
		return NewImagesModel(
			args, flags, subCommand,
		)
	},
)

func init() {
	describeCmd.Flags().StringP("params", "p", "params.yaml", "Path to the build parameters file")
}
