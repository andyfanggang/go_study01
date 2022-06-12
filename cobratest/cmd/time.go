package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"wm-motor.com/Infra/cobratest/internal/timer"
)

var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式转换",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果:%s", nowTime.Format("2006-01-02 15:04:05"))
	},
}

var CalculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需要时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		t, err := timer.GetCalculateTime(nowTime, duration)
		if err != nil {
			log.Fatalf("计算错误!")
		}
		log.Printf("输出结果:%s", t.Format("2006-01-02 15:04:05"))
	},
}

func init() {
	CalculateTimeCmd.Flags().StringVarP(&duration, "dur", "d", "", "请输入时间差")
	timeCmd.AddCommand(nowTimeCmd) //将now命令基加到time命令
	timeCmd.AddCommand(CalculateTimeCmd)

}
