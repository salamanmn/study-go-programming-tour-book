package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-programming-tour/tour/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := time.Now()
		log.Printf("输出结果：%s, %d", nowTime.Format(time.RFC3339), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var calculateTimer time.Time
		var layout string
		if calculateTime == "" {
			calculateTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			} 
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			calculateTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				log.Fatalln("error!")
				t, _ := strconv.Atoi(calculateTime) //人为保证第三种格式是时间戳格式
				calculateTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(calculateTimer, duration)
		
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果： %s %d", t.Format(layout), t.Unix())

	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位是时间戳或格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效时间单位是ns,us,ms等")
}
