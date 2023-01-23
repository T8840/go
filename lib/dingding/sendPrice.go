package dingding

import (
	"fmt"
	ddMsg "godemo/lib/dingding/dingdingNotification"
)

var WEB_HOOK = ""

func SendBTCPrice(){
	/*
	//text 格式发送，注意关键词
	send_text := ddMsg.DingMSG_text("BTPrice: 230000", true)
	resp := ddMsg.DingdingMSG(WEB_HOOK, send_text)
	fmt.Println(resp.Status)
	*/

	//markdown 格式发送，注意关键词，关键词放到标题或内容里都可以
	content := fmt.Sprintf("# BTPrice \n> #### BTPrice: 230000")
	send_markdown := ddMsg.DingMSG_markdown("BTPrice", content, true)
	resp1 := ddMsg.DingdingMSG(WEB_HOOK, send_markdown)
	fmt.Println(resp1.Status)
}