package cmd

import (
	"github.com/spf13/cobra"
	"simpledocker/container"
)

var RunCmd = &cobra.Command{
	Use:              "run [flags] [COMMAND]",
	Short:            "Run a command in a new container",
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		container.RunParam.Id = container.GenContainerId(32)
		if container.RunParam.Name == "" {
			container.RunParam.Name = container.RunParam.Id
		}
		container.RunParam.Image = args[0]
		container.RunParam.Cmd = args[1:]
		container.Run()
	},
}

func init() {
	RunCmd.Flags().BoolVar(&container.RunParam.TTY, "tty", false, "enable tty")
	RunCmd.Flags().BoolVarP(&container.RunParam.Detach, "detach", "d", false, "Run container in background and print container ID")
	RunCmd.Flags().StringVarP(&container.RunParam.Memory, "memory", "m", "", "memory limit")
	RunCmd.Flags().StringVar(&container.RunParam.Name, "name", "", "container name")
	RunCmd.Flags().StringSliceVarP(&container.RunParam.Env, "env", "e", nil, "container env")
	RunCmd.Flags().StringSliceVarP(&container.RunParam.Volumes, "volume", "v", nil, "container volume")
	RunCmd.Flags().StringVarP(&container.RunParam.Port, "port", "p", "", "container port")
}
