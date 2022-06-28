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
		container.ProcessRunParam.Id = container.GenContainerId(32)
		if container.ProcessRunParam.Name == "" {
			container.ProcessRunParam.Name = container.ProcessRunParam.Id
		}
		container.ProcessRunParam.Image = args[0]
		container.ProcessRunParam.Cmd = args[1:]
		container.Run()
	},
}

func init() {
	RunCmd.Flags().BoolVar(&container.ProcessRunParam.TTY, "tty", false, "enable tty")
	RunCmd.Flags().BoolVarP(&container.ProcessRunParam.Detach, "detach", "d", false, "Run container in background and print container ID")
	RunCmd.Flags().StringVarP(&container.ProcessRunParam.Memory, "memory", "m", "", "memory limit")
	RunCmd.Flags().StringVar(&container.ProcessRunParam.Name, "name", "", "container name")
	RunCmd.Flags().StringSliceVarP(&container.ProcessRunParam.Env, "env", "e", nil, "container env")
	RunCmd.Flags().StringSliceVarP(&container.ProcessRunParam.Volumes, "volume", "v", nil, "container volume")
	RunCmd.Flags().StringVarP(&container.ProcessRunParam.Port, "port", "p", "", "container port")
}
