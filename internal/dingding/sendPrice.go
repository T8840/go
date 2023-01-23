package dingding

import (
	"fmt"
	"time"
	ddMsg "godemo/lib/dingding/dingdingNotification"
	conf "godemo/utils/conf"
	"github.com/robfig/cron"
	goexv2 "github.com/nntaoli-project/goex/v2"
	// "github.com/nntaoli-project/goex/v2/logger"
	"github.com/nntaoli-project/goex/v2/model"
	// "github.com/nntaoli-project/goex/v2/options"
	
)


func SendPrice(){
	c := conf.ViperGetConf("dingding")
	group := c.(map[string]interface{})  
	WEB_HOOK := group["web_hook"].(string)


	//text 格式发送，注意关键词
	// send_text := ddMsg.DingMSG_text("BTPrice: 230000", true)
	// resp := ddMsg.DingdingMSG(WEB_HOOK, send_text)
	// fmt.Println(resp.Status)
	
	//markdown 格式发送，注意关键词，关键词放到标题或内容里都可以
	
	cronNew := cron.New()
	spec := "*/10 * * * *"
	cronNew.AddFunc(spec, func() {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		btcTick, _, _ := goexv2.OKx.Futures.GetTicker(model.CurrencyPair{Symbol: "BTC-USDT"})
		aptTick, _, _ := goexv2.OKx.Futures.GetTicker(model.CurrencyPair{Symbol: "APT-USDT"})
		dogeTick, _, _ := goexv2.OKx.Futures.GetTicker(model.CurrencyPair{Symbol: "DOGE-USDT"})

		btcPrice := btcTick.Last
		dogePrice := dogeTick.Last
		aptPrice := aptTick.Last

		content := fmt.Sprintf("# Price %s \n> ### BTCPrice: %f \n> ### APTPrice: %f \n> ### DogePrice: %f",nowTime,btcPrice,aptPrice,dogePrice)
		send_markdown := ddMsg.DingMSG_markdown("BTPrice", content, true)
		resp1 := ddMsg.DingdingMSG(WEB_HOOK, send_markdown)
		fmt.Println(resp1.Status)
	})
	cronNew.Start()
	select {}
	

}