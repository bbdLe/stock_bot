package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stock_bot/internal/log"
	"stock_bot/internal/util"
	"time"
)

type StockDayData struct {
	Bbd 	string `json:"bbd"`
	Cjl 	string `json:"cjl"`
	Ddc 	string `json:"ddc"`
	Ddx 	string `json:"ddx"`
	Ddx10 	string `json:"ddx10"`
	Ddx3 	string `json:"ddx3"`
	Ddx5 	string `json:"ddx5"`
	Ddx60 	string `json:"ddx60"`
	Ddy 	string `json:"ddy"`
	Dsb 	string `json:"dsb"`
	Dtime 	string `json:"dtime"`
	KaiPan 	string `json:"kaipan"`
	Spj 	string `json:"spj"`
	Tcl 	string `json:"tcl"`
	Tdc 	string `json:"tdc"`
	Xdc 	string `json:"xdc"`
	Zdc 	string `json:"zdc"`
	Zf 		string `json:"zf"`
	ZuiDi 	string `json:"zuidi"`
	ZuiGao 	string `json:"zuidi"`
}

func (self *StockDayData) String() string {
	msg := ""

	msg += fmt.Sprintf("股价 : %s ", util.ConvVal2MarkDown(self.Spj))
	msg += fmt.Sprintf("涨幅 : %s ", util.ConvPercent2MarkDown(self.Zf))
	msg += fmt.Sprintf("通吃率 : %s ", util.ConvPercent2MarkDown(self.Tcl))
	msg += fmt.Sprintf("ddx : %s ", util.ConvVal2MarkDown(self.Ddx))
	msg += fmt.Sprintf("ddx3 : %s ", util.ConvVal2MarkDown(self.Ddx3))
	msg += fmt.Sprintf("ddy : %s ", util.ConvVal2MarkDown(self.Ddy))
	msg += fmt.Sprintf("小单差 : %s ", util.ConvVal2MarkDown(self.Xdc))
	msg += fmt.Sprintf("中单差 : %s ", util.ConvVal2MarkDown(self.Zdc))
	msg += fmt.Sprintf("大单差 : %s ", util.ConvVal2MarkDown(self.Tdc))
	msg += fmt.Sprintf("特大单差 : %s ", util.ConvVal2MarkDown(self.Tdc))
	msg += fmt.Sprintf("成交量 : %s万", util.ConvVal2MarkDown(self.Cjl))

	return msg
}

type Reply struct {
	Data 		[]*StockDayData `json:"data"`
	UpdateTime 	string `json:"updatetime"`
}

const (
	dataUrlTemp = "http://ddx.chaguwang.cn/ddelist.php?code=%s&bdate=%s&edate=%s"
	DateFormat  = "2006-01-02"
	DateTimeFormat  = "2006-01-02 15:04:05"
)

func GetStock(stockId string, begin time.Time, end time.Time) (*Reply, error) {
	beginStr := begin.Format(DateFormat)
	endStr   := end.Format(DateFormat)
	url 	 := fmt.Sprintf(dataUrlTemp, stockId, beginStr, endStr)
	log.Logger.Info(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reply := new(Reply)
	data, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &reply); err != nil {
		log.Logger.Error("json unmarshal fail")
		return nil, err
	}

	return reply, nil
}
