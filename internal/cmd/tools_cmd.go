package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	logger "godemo/utils/logger"
	// db "godemo/internal/db"
	dd "godemo/internal/dingding"

)

var (
	tools string
	type_name string

)

var scriptsCmd = &cobra.Command{
	Use:   "t",
	Short: "command for common tools",
	Long: ` command for common tools:
	1、./goctl t -h
	2、./goctl t -t db -n mysql
		-t: 工具类型 db/dingding/conf
			db: db工具类型
				-n:	mysql/mongo
			dingding: dingding工具类型
				-n: btc/apt/doge
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pid:",os.Getpid())
		var desc = "tools_" + strconv.Itoa(os.Getpid())
		logger.SetLogFile(desc)
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch tools {
			
		case "db":
			logger.Logger.Println("db")
			switch type_name {
			case  "mysql":
				logger.Logger.Println("mysql")
				break
			case  "mongo","mongodb":
				logger.Logger.Println("mongo")
				// db.Insert()
				// db.InsertToStudent()
			default:
				fmt.Println("conf")
			}
		case "dd","dingding":
			logger.Logger.Println("---Use dingding Tools---")
			switch type_name {
			case  "doge":
				logger.Logger.Println("doge")
				break
			case  "apt":
				logger.Logger.Println("apt")
			default:
				logger.Logger.Println("dingding sendPrice")
				dd.SendPrice()
			}
		case "env":
			break
		
		default:
			fmt.Printf("请输入合适的tool类型：log_tool/blade/conf/env_check/")
		
	}

	},
}

func init() {
	rootCmd.AddCommand(scriptsCmd)
	scriptsCmd.Flags().StringVarP(&tools, "tools", "t", "", "tool name")
	scriptsCmd.Flags().StringVarP(&type_name, "type_name", "n", "", "type name")

}