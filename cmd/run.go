package cmd

import (
	"github.com/spf13/cobra"
	"simpledocker/container"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command in a new container",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []container.Option
		if flags.tty {
			opts = append(opts, container.SetTTY(true))
		}
		name := flags.name
		if len(name) == 0 {
			name = container.GenContainerId(32)
		}
		opts = append(opts, container.SetContainerName(name))

		if len(flags.env) > 0 {
			opts = append(opts, container.SetEnvs(flags.env))
		}
		process := container.NewProcess(opts...)
		process.Create()

	},
}
var flags struct {
	tty    bool
	detach bool
	memory string
	name   string
	env    []string
	port   string
}

func init() {
	RunCmd.Flags().BoolVarP(&flags.tty, "tty", "t", false, "enable tty")
	RunCmd.Flags().BoolVarP(&flags.detach, "detach", "d", false, "Run container in background and print container ID")
	RunCmd.Flags().StringVarP(&flags.memory, "memory", "m", "", "memory limit")
	RunCmd.Flags().StringVar(&flags.name, "name", "", "container name")
	RunCmd.Flags().StringSliceVarP(&flags.env, "env", "e", nil, "container env")
	RunCmd.Flags().StringVarP(&flags.port, "port", "p", "", "container port")
}
