package conf

import (
	"fmt"
	"log"
	"os"
	"strings"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	oscmd "godemo/utils/oscmd"
)

func uri2map(uri string) (map[string]string, error) {
	m := make(map[string]string)
	if len(uri) < 1 { // 空字符串
		return m, errors.New("uri is none")
	}
	//if uri[0:1] == "?" { // 有没有包含？,有的话忽略。
	//	uri = uri[1:]
	//}
	if strings.Contains(uri,"?") {
		index:= strings.Index(uri,"?") // 有没有包含？,有的话截取后面的字符串
		uri = uri[index+1:]
	}
	// 去除特殊字符
	uri = strings.Replace(uri,"'","",-1)
	uri = strings.Replace(uri,"\"","",-1)
	// http://10.110.41.214:8848/nacos/v1/cs/configs?dataId=dtssdk_cpp_test_conf&group=DEFAULT_GROUP&tenant=9424bdb7-9ad4-4675-afc8-b8a9b53cf45c
	// 首先分解&（sessionid=22222&token=3333 )变成 sessionid=222222 和 token=3333
	pars := strings.Split(uri, "&")
	for _, par := range pars {
		// 然后分解 sessionid=222222 和 token=3333
		parkv := strings.Split(par, "=")
		m[parkv[0]] = parkv[1] // 等号前面是key,后面是value
	}
	log.Println(m)
	return m, nil
}

func write2file(file string,content string) {
	dstFile,err :=os.Create(file)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(content)
}

func GetNacosConf(url string) string{

	// url解析group/dataId/NamespaceId
	m, _ :=uri2map(url)
	fmt.Println("【UseNacosConf】group:" + m["group"] + ", dataId:" + m["dataId"] + ", tenant:" + m["tenant"])
	sc := []constant.ServerConfig{{
		IpAddr: "10.110.41.214",
		Port:   8848,
	}}

	cc := constant.ClientConfig{
		NamespaceId:         m["tenant"], // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: m["dataId"],
		Group:  m["group"],
	})

	if err != nil {
		panic(err)
	}
	//fmt.Println("Get Nacos Conf:" + content) //字符串 - yaml
	project_path := oscmd.GetProjectPath()
	file_name := project_path + "/conf/nacos_conf/" + m["dataId"] +".yaml"
	write2file(file_name,content)
	return file_name

	// //持续监听nacos配置变动
	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: "user-web.yaml",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生了变化...")
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//	},
	//})
	//
}