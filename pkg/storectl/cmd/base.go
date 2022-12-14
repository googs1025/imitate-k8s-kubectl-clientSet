package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"log"
	"practice_ctl/pkg/storectl/config"
)

type CmdMetaData struct {
	Use string
	Short string
	Example string
}

var storeCmdMetaData *CmdMetaData
func init() {
	storeCmdMetaData = &CmdMetaData{
		Use: "storeclt [flags]",
		Short: "模仿kubectl",
		Example: "storeclt [flags]",
	}

}

// TODO: apple car命令行

// RunCmd
func RunCmd() {
	cmd := &cobra.Command{
		Use:          storeCmdMetaData.Use,
		Short:        storeCmdMetaData.Short,
		Example:      storeCmdMetaData.Example,
		SilenceUsage: true,
	}

	configRes := config.LoadConfigFile()

	// list操作
	listCmd := ListCommand(configRes)
	createCmd := CreateCommand(configRes)

	//加入子命令
	cmd.AddCommand(versionCmd, listCmd, createCmd)
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}


var cfgFlags *genericclioptions.ConfigFlags

func MergeFlags(cmds ...*cobra.Command) {
	cfgFlags = genericclioptions.NewConfigFlags(true)
	for _, cmd := range cmds {
		cfgFlags.AddFlags(cmd.Flags())
	}

}