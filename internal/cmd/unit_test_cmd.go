package cmd

import (
	"os/signal"
	"strings"
	"syscall"
	//"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	oscmd "godemo/utils/oscmd"
	conf "godemo/utils/conf"
	logger "godemo/utils/logger"
	"strconv"
)

var (
	web_test_port string
	test_case string
	test_conf string
	exec_cmd *exec.Cmd
	out_report_path string
)

var unitTestCmd = &cobra.Command{
	Use:   "u",
	Short: "command for unit test",
	Long: `command for unit test:
	1、./goctl u -h
		打印帮助信息
	2、Go Test
     使用本地conf配置：
     ./goctl u -t a -v 2.1.6   # 未指定-l为go test
     ./goctl u -t a/c/p/i -o/--out test.html
          注释：-t 参数后面可支持特定类型case：all/exception/consumer/producer
     ./goctl u -t w -p 9001
		  注释：-w 开启goconvey web页面
     使用Nacos配置：
     
    3、Cpp Test
     使用本地conf配置：./goctl u -t a -l cpp -v 2.1.6
     使用Nacos配置：  ./goctl u -t a -l cpp -v 2.1.6 -f 'http://10.110.41.214:8848/nacos/v1/cs/configs?dataId=dtssdk_cplus_2.1.6&group=DEFAULT_GROUP&tenant=9424bdb7-9ad4-4675-afc8-b8a9b53cf45c' 
    4、Java Test
     使用本地conf配置：./goctl u -t a -l java -v 2.1.6
     使用Nacos配置：  ./goctl u -t a -l java -v 2.1.6 -f 'http://10.110.41.214:8848/nacos/v1/cs/configs?dataId=dtssdk_java_2.1.6&group=DEFAULT_GROUP&tenant=9424bdb7-9ad4-4675-afc8-b8a9b53cf45c' 

	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pid:",os.Getpid())
		var desc = "unit_test_" + strconv.Itoa(os.Getpid())
		logger.SetLogFile(desc)
		project_path := oscmd.GetProjectPath()
		switch language {
		case "j", "java":
			fmt.Println("[PreRun]:Java")
			java_test_dir := project_path + "/java_test/"
			fmt.Println(java_test_dir)
			// cur_branch, err := oscmd.RunGitCmd(java_test_dir, "git", "symbolic-ref", "--short", "HEAD")
			// if err != nil {
			// 	fmt.Println("git branch err:", err)
			// }
			// fmt.Println("cur_branch:", cur_branch)

			// switch test_version {
			// case "2.1.8":
			// 	res, err := oscmd.RunGitCmd(java_test_dir, "git", "checkout", "v2.1.8")
			// 	if err != nil {
			// 		fmt.Println("git checkout err:", err)
			// 	}
			// 	fmt.Println("git checkout res:", res)
			// default:
			// 	res, err := oscmd.RunGitCmd(java_test_dir, "git", "checkout", "master")
			// 	if err != nil {
			// 		fmt.Println("git checkout err:", err)
			// 	}
			// 	fmt.Println("git checkout res:", res)
			// }
			// git_pull_res,err := oscmd.RunGitCmd(java_test_dir, "git", "pull")
			// if err!=nil{
			// 	fmt.Println("git pull err:",err)
			// }
			// fmt.Println("git pull res:",git_pull_res)
			

		case "c", "cpp", "cplus":
			fmt.Println("[PreRun]:Cpp")
			cpp_test_dir := project_path + "/cplus_test/"
			fmt.Println(cpp_test_dir)
			// cur_branch,err := oscmd.RunGitCmd(cpp_test_dir, "git", "symbolic-ref","--short", "HEAD")
			// if err!=nil{
			// 	fmt.Println("git branch err:",err)
			// }
			// fmt.Println("cur_branch:",cur_branch)

			// switch test_version {
			// case "2.1.6" :
			// 	res,err := oscmd.RunGitCmd(cpp_test_dir, "git", "checkout","v2.1.6")
			// 	if err!=nil{
			// 		fmt.Println("git checkout err:",err)
			// 	}
			// 	fmt.Println("git checkout res:",res)

			// default:
			// 	res,err := oscmd.RunGitCmd(cpp_test_dir, "git", "checkout","master")
			// 	if err!=nil{
			// 		fmt.Println("git checkout err:",err)
			// 	}
			// 	fmt.Println("git checkout res:",res)
			// }
			// git_pull_res,err := oscmd.RunGitCmd(cpp_test_dir, "git", "pull")
			// if err!=nil{
			// 	fmt.Println("git pull err:",err)
			// }
			// fmt.Println("git pull res:",git_pull_res)
		default:
			fmt.Println("[PreRun]:Go")
			go_test_dir := project_path + "/go_test/"
			fmt.Println(go_test_dir)
			// 	cur_branch,err := oscmd.RunGitCmd(go_test_dir, "git", "symbolic-ref","--short", "HEAD")
			// 	if err!=nil{
			// 		fmt.Println("git branch err:",err)
			// 	}
			// 	fmt.Println("cur_branch:",cur_branch)

			// 	switch test_version {
			// 	case "2.1.6" :
			// 		res,err := oscmd.RunGitCmd(go_test_dir, "git", "checkout","master_dev_1.2.1")
			// 		if err!=nil{
			// 			fmt.Println("git checkout err:",err)
			// 		}
			// 		fmt.Println("git checkout res:",res)

			// 	default:
			// 		res,err := oscmd.RunGitCmd(go_test_dir, "git", "checkout","v1.2.1")
			// 		if err!=nil{
			// 			fmt.Println("git checkout err:",err)
			// 		}
			// 		fmt.Println("git checkout res:",res)
			// 	}
			// 	git_pull_res,err := oscmd.RunGitCmd(go_test_dir, "git", "pull")
			// 	if err!=nil{
			// 		fmt.Println("git pull err:",err)
			// 	}
			// 	fmt.Println("git pull res:",git_pull_res)
			// 	rm_res,err :=oscmd.CmdAndChangeDir(go_test_dir,"rm", []string{ "-f","go.mod","go.sum"})
			// 	fmt.Println("[PreRun]:RM go.mod go.sum:",rm_res,err)
			}

		fmt.Println("[PreRun]:Done!")
	},
	Run: func(cmd *cobra.Command, args []string) {
		// var test_dir string
		project_path := oscmd.GetProjectPath()
		switch language {
		case "j","java":
			switch test_case {
			case "a","all":
				fmt.Println("java test")
				// java_test_dir := project_path + "/java-test/"
				// cp_res,err :=oscmd.CmdAndChangeDir(project_path,"cp", []string{ "coyote",java_test_dir})
				// fmt.Println("[Run]:CP Coyote",cp_res,err)
				// mvn_res,err :=oscmd.CmdAndChangeDir(java_test_dir,"mvn", []string{ "clean","package","-Dmaven.test.skip=true"})
				// fmt.Println("[Run]:JavaMvn",mvn_res,err)
				// report_path := out_report_path + "/index.html"
				// if strings.Contains(test_conf,"http") {
				// 	test_conf_yaml := utils.GetNacosConf(test_conf)
				// 	res :=  utils.WatchExecCmdDir(project_path, "./coyote", []string{"-c" ,test_conf_yaml, "-out" ,report_path,"-timeout", "40m"})
				// 	fmt.Println(res)
				// } else {
				// 	test_conf_yaml := project_path + "/conf/" + "java_test_" + test_version + ".yaml"
				// 	res := utils.WatchExecCmdDir(java_test_dir, "./coyote", []string{"-c", test_conf_yaml, "-out", report_path, "-timeout", "40m"})
				// 	fmt.Println(res)
				// 	}
				
			default:
				{}
			}
		case "c","cpp","cplus":
			cpp_test_dir := project_path + "/cplus_test"
			fmt.Println("cpp_test_dir:",cpp_test_dir)
			switch test_case {
			case "a","all":
				report_path := out_report_path + "/index.html"
				if strings.Contains(test_conf,"http") {
					test_conf_yaml := conf.GetNacosConf(test_conf)
					res :=  oscmd.WatchExecCmdDir(project_path, "./coyote", []string{"-c" ,test_conf_yaml, "-out" ,report_path,"-timeout", "40m"})
					fmt.Println(res)
				} else {
					test_conf_yaml := "conf/" + "cpp_test_" + test_version + ".yaml"
					res := oscmd.WatchExecCmdDir(project_path, "./coyote", []string{"-c", test_conf_yaml, "-out", report_path, "-timeout", "40m"})
					fmt.Println(res)
				}
			default:
					{
					}
				}
		default:
			go_test_dir := project_path + "/internal/go_test"
			fmt.Println("go_test_dir:",go_test_dir)
			// 获取关于go_test配置
			//go_test_conf := utils.ViperGetConf("go_test")
			//fmt.Println(go_test_conf)
			switch test_case {
			case "w","web":
				res,_ := oscmd.CmdAndChangeDir(go_test_dir, "goconvey", []string{"-host","0.0.0.0","-port",web_test_port})
				fmt.Println(res)
				c := make(chan os.Signal)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				<-c
			case "a","all":
				report_path := out_report_path + "/index.html"
				if strings.Contains(test_conf,"http") {
					test_conf_yaml := conf.GetNacosConf(test_conf)
					res :=  oscmd.WatchExecCmdDir(project_path, "./coyote", []string{"-c" ,test_conf_yaml, "-out" ,report_path,"-timeout", "40m"})
					fmt.Println(res)
				} else {
					test_conf_yaml := "conf/" + "go_test_" + test_version + ".yaml"
					res :=  oscmd.WatchExecCmdDir(project_path, "./coyote", []string{"-c" ,test_conf_yaml, "-out" ,report_path,"-timeout", "40m"})
					fmt.Println(res)
				}

			default:
				{

				}
			
			}
		
	}
	},
}

func init() {
	rootCmd.AddCommand(unitTestCmd)
	unitTestCmd.Flags().StringVarP(&out_report_path, "out", "o", "/Users/greyhead/Documents/Go/godemo/report", "output report path")
	unitTestCmd.Flags().StringVarP(&test_case, "test_case", "t", "", "go test special case: all/producer/consumer/interface")
	unitTestCmd.Flags().StringVarP(&test_conf, "test_conf", "f", "local", "test conf from local file or from nacos")
	unitTestCmd.Flags().StringVarP(&web_test_port, "web_test_port", "p","8080" , "goconvey web port")
	unitTestCmd.Flags().StringVarP(&language, "language", "l", "go", "unit test language")
	unitTestCmd.Flags().StringVarP(&test_version, "test_version", "v", "1.0", "unit test version")

}