package test

import (
	"fmt"
	conf "godemo/utils/conf"
	"testing"
	oscmd "godemo/utils/oscmd"

)
   
func TestConf(t *testing.T) {
	local_path:= oscmd.GetLocalDirPath()
	fmt.Println(local_path)
	// For Windows
	conf_path := local_path + "godemo\\conf\\conf.yaml"
	var c conf.DingDingConf
	m := c.GetDingDingConf(conf_path)
	fmt.Println(m)
}

func TestDingDing(t *testing.T) {
	c := conf.ViperGetConf("dingding")
	group := c.(map[string]interface{})  
	fmt.Println(group["web_hook"])
}
