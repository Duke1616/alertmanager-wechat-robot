package initial

import (
	"github.com/spf13/cobra"
)

var (
	debug bool
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "alertmanager-wechat-robot 服务初始化",
	Long:  "alertmanager-wechat-robot 服务初始化",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		exec, err := NewExecutorFromCLI()
		cobra.CheckErr(err)

		// 初始化管理员用户
		err = exec.InitTarget(ctx)
		cobra.CheckErr(err)

		// 初始化报警群组
		err = exec.InitAdminUser(ctx)
		cobra.CheckErr(err)
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug info")
}
