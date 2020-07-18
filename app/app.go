package app

import (
	"fmt"
	"os"
	"time"

	"stock_bot/internal/log"
	"stock_bot/internal/util"

	"github.com/bastengao/chinese-holidays-go/holidays"
	"github.com/happysooner/WechatWorkRobot"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	robot *WechatWorkRobot.Robot
)

const (
	msgTemp = "【%s】 %s<br/>"
)

func Run() {
	stockBot := cli.NewApp()
	stockBot.Name = "stock bot"
	stockBot.Usage = "bot for stock"
	stockBot.Before = Init
	stockBot.Action = Action

	if err := stockBot.Run(os.Args); err != nil {
		log.Logger.Error("appApp exit", zap.Error(err))
	}
}

func Init(ctx *cli.Context) error {
	if err := util.WritePid("../etc/robot.pid"); err != nil {
		return err
	}

	if err := initCfg("../etc/conf.toml"); err != nil {
		return err
	}
	robot = &WechatWorkRobot.Robot{Key: GetConfig().EnvConfig.WebHookKey}

	return nil
}

func Action(ctx *cli.Context) error {
	for {
		RunData()
		time.Sleep(time.Minute * 10)
	}
}

func RunData() {
	//跳过节假日
	ok, err := holidays.IsHoliday(time.Now())
	if err != nil {
		log.Logger.Error("check work day fail", zap.Error(err))
		return
	}
	if ok {
		return
	}

	//是否市场开启时间
	if !util.IsMarkTime() {
		return
	}

	m, err := FetchStockData()
	if err != nil {
		log.Logger.Error("fetch stock data fail", zap.Error(err))
		return
	}

	SendData(GetConfig().EnvConfig.StockList, m)

	return
}

func SendData(orderList []Stock, m map[string]*StockDayData) {
	if len(m) <= 0 {
		return
	}

	msg := "更新时间:" + time.Now().Format(DateTimeFormat) + "<br/>"
	count := 0

	for _, stock := range orderList {
		v, ok := m[stock.Name]
		if !ok {
			continue
		}

		data := fmt.Sprintf(msgTemp, stock.Name, v.String())
		msg += data

		count += 1

		// 分批发送
		if count > 6 {
			count = 0
			msg = msg[: len(msg) - 5]

			resp, err := robot.SendMarkdown(msg)
			if err != nil || resp.ErrorCode != 0 {
				log.Logger.Error("send fail", zap.Error(err), zap.String("msg", resp.ErrorMessage))
			}

			msg = ""
		}
	}

	if len(msg) > 0 {
		msg = msg[: len(msg) - 5]

		resp, err := robot.SendMarkdown(msg)
		if err != nil || resp.ErrorCode != 0 {
			log.Logger.Error("send fail", zap.Error(err), zap.String("msg", resp.ErrorMessage))
		}
	}
}

func FetchStockData() (map[string]*StockDayData, error) {
	cfg := GetConfig()

	m := make(map[string]*StockDayData)
	stockList := cfg.EnvConfig.StockList
	for _, stock := range stockList {
		reply, err := GetStock(stock.Id, time.Now().Add(time.Hour * -24), time.Now().Add(time.Hour * -24))
		if err != nil {
			log.Logger.Error("get stock fail", zap.Error(err))
			continue
		}
		if len(reply.Data) < 1 {
			log.Logger.Error("stock data is empty", zap.String("id", stock.Id))
			continue
		}

		m[stock.Name] = reply.Data[0]
	}

	return m, nil
}