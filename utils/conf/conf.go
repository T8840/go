package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	yaml "gopkg.in/yaml.v3"
	"reflect"
	oscmd "godemo/utils/oscmd"
)

type GoTestConf struct {
	Topic struct {
		CommonTopic string `yaml:"common_topic_name"`
		RandomTopic string `yaml:"random_topic_name"`
	}

}

func (c *GoTestConf) GetConf(path string) *GoTestConf {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

type DingDingConf struct {
	Group struct {
		GroupName string `yaml:"group_name"`
		WebHook string `yaml:"web_hook"`
	}

}


func (c *DingDingConf) GetDingDingConf(path string) *DingDingConf {
	
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func ViperGetConf( key string)  interface{}{
	project_path := oscmd.GetLocalDirPath()
	// Linux
	//conf_path := project_path + "godemo/conf/"
	// Windows
	conf_path := project_path + "godemo\\conf\\"

	viper.Set("fileDir",conf_path)
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(conf_path)
	err :=viper.ReadInConfig()
	if err !=nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("key:" , key ,viper.Get(key))
	return viper.Get(key)
}


func DtssdkGetConf(path string)  interface{}{
	Config := viper.New()
	viper.SetConfigName("dtssdk_config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(path)

	if err := Config.ReadInConfig();err!= nil {
		if _,ok :=err.(viper.ConfigFileNotFoundError);ok {
			fmt.Println("找不到配置dtssdk文件")
		} else {
			fmt.Println("dtssdk配置文件出错")
		}
	}
	Config.WatchConfig() // 监控配置文件并且热加载程序，不重启就可以加载新的配置文件

	Config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("Dtssdk Config file Changed: %s",in.Name)
	})
	// TODO:合理的解析未做，目前无法打印
	if reflect.TypeOf(Config).Kind() == reflect.Slice {
		s := reflect.ValueOf(Config)
		for i :=0; i< s.Len(); i++ {
			fmt.Println(s.Index(i),"=",s,"\n")
		}
	}
	return Config

}


func DtssdkGetKeyConf( key string)  interface{}{
	dtssdk_conf_path := "/data/dtssdk-sdk/conf/"
	viper.Set("fileDir",dtssdk_conf_path)
	viper.SetConfigName("dtssdk_config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(dtssdk_conf_path)
	err :=viper.ReadInConfig()
	if err !=nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("key:" , key ,viper.Get(key))
	return viper.Get(key)
}


func print_json(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}