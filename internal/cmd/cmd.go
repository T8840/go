package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	oscmd "godemo/utils/oscmd"
)

var (
	cfgFile string
	language string
	model string
	test_version string

)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "cmd",
	Long: `Home Cmd.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

// 执行 rootCmd 命令并检测错误
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// 加载运行初始化配置
	cobra.OnInitialize(initConfig)
	// rootCmd，命令行下读取配置文件，持久化的 flag，全局的配置文件
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $project/conf/conf.yaml)")
	// local flag，本地化的配置
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// 初始化配置的一些设置
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // viper 设置配置文件
	} else {// 上面没有指定配置文件，下面就读取 conf 下的 conf.yaml文件
		// 配置文件参数设置
		//home, err := os.UserHomeDir()
		//cobra.CheckErr(err)
		//viper.AddConfigPath(home)
		project_path := oscmd.GetProjectPath()
		conf_dir := project_path + "/conf"
		viper.AddConfigPath(conf_dir)
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {// 读取配置文件
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}