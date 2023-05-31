package devfile

import (
	"github.com/redhat-developer/alizer/go/pkg/apis/recognizer"
	"github.com/redhat-developer/alizer/go/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	registry   string
	minVersion string
	maxVersion string
	logLevel   string
)

func NewCmdDevfile() *cobra.Command {
	devfileCmd := &cobra.Command{
		Use:   "devfile",
		Short: "Select a devfile from a list of existing devfiles (from a devfile registry or other storage) based on the information extracted from the source tree.",
		Long:  "Select a devfile from a list of existing devfiles (from a devfile registry or other storage) based on the information extracted from the source tree.",
		Args:  cobra.MaximumNArgs(4),
		Run:   doSelectDevfile,
	}
	devfileCmd.Flags().StringVarP(&registry, "registry", "r", "", "registry where to download the devfiles. Default value: https://registry.devfile.io")
	devfileCmd.Flags().StringVarP(&minVersion, "min-version", "", "", "minimum version of devfile schemaVersion. Default value: nil")
	devfileCmd.Flags().StringVarP(&maxVersion, "max-version", "", "", "maximum version of devfile schemaVersion. Default value: nil")
	devfileCmd.Flags().StringVarP(&logLevel, "log", "", "", "log level for alizer. Default value: error. Accepted values: [debug, info, warning]")
	return devfileCmd
}

func doSelectDevfile(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		utils.PrintNoArgsWarningMessage(cmd.Name())
		return
	}
	if registry == "" {
		registry = "https://registry.devfile.io/index"
	}
	utils.PrintPrettifyOutput(recognizer.SelectDevFilesFromRegistry(args[0], registry))
}
