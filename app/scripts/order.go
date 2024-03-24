package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"time"

	"log"
)

var token string

func main() {
	cli := resty.New()
	var result = map[string]interface{}{}
	t := time.NewTimer(time.Second * 4)
	c := getPrice()
	//下单脚本
	for {
		select {
		case price := <-c:
			if token == "" {
				continue
			}
			asksLevel, bidsLevel, err := getDepthLevel()
			if err != nil {
				log.Printf("获取深度失败，%v", err)
			}
			if asksLevel < 15 {
				//卖盘 价差3%
				price = price * 1.03
				_, err = cli.R().
					SetHeader("gexToken", token).
					SetHeader("Content-Type", "application/json").
					SetBody(map[string]interface{}{
						"symbol_id":   14,
						"symbol_name": "BTC_USDT",
						"price":       cast.ToString(price),
						"qty":         "1",
						"amount":      "",
						"side":        2,
						"order_type":  2,
					}).
					SetResult(&result).
					Post("http://api.gex.com/order/v1/create_order")
				if err != nil {
					log.Printf("下单失败 %v", err)
				}
				code, ok := result["code"]
				if !ok || cast.ToInt64(code) != 0 {
					log.Printf("下单失败 %v", result)
					continue
				}
				log.Println("下单成功", price)
			}
			if bidsLevel < 15 {
				//卖盘 价差3%
				price = price * 0.97
				_, err = cli.R().
					SetHeader("gexToken", token).
					SetHeader("Content-Type", "application/json").
					SetBody(map[string]interface{}{
						"symbol_id":   14,
						"symbol_name": "BTC_USDT",
						"price":       cast.ToString(price),
						"qty":         "1",
						"amount":      "",
						"side":        1,
						"order_type":  2,
					}).
					SetResult(&result).
					Post("http://api.gex.com/order/v1/create_order")
				if err != nil {
					log.Printf("下单失败 %v", err)
				}
				code, ok := result["code"]
				if !ok || cast.ToInt64(code) != 0 {
					log.Printf("下单失败 %v", result)
					continue
				}
				log.Println("下单成功", price)
			}

		case <-t.C:
			login()
		}
	}
}
func getDepthLevel() (int, int, error) {
	cli := resty.New()
	var result = map[string]interface{}{}
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"symbol":"BTC_USDT", "level":300}`).
		SetResult(&result).
		Post("http://api.gex.com/quotes/v1/get_depth_list")
	if err != nil {
		log.Printf("获取深度失败 %v", err)
		return 0, 0, err
	}
	code, ok := result["code"]
	if !ok || cast.ToInt64(code) != 0 {
		log.Printf("获取深度失败 %v", result)
		return 0, 0, err
	}
	d := result["data"].(map[string]interface{})
	bids := d["bids"].([]interface{})
	asks := d["asks"].([]interface{})
	return len(asks), len(bids), nil
}
func login() {
	cli := resty.New()
	var result = map[string]interface{}{}
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"lisi", "password":"lisilisi"}`).
		SetResult(&result).
		Post("http://api.gex.com/account/v1/login")
	if err != nil {
		log.Panicf("err %v", err)
	}
	code, ok := result["code"]
	if !ok {
		log.Printf("data=%v", result)
	}
	if code != 0 {
		log.Printf("data=%v", result)
	}
	data := result["data"].(map[string]interface{})
	token = cast.ToString(data["token"])

}

type PriceData struct {
	C interface{} `json:"c"`
	P float64     `json:"p"`
	S string      `json:"s"`
	T int64       `json:"t"`
	V float64     `json:"v"`
}
type Resp struct {
	Data []*PriceData `json:"data"`
}

func getPrice() <-chan float64 {
	f := make(chan float64)
	c, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token=cimf8epr01qlsedscmvgcimf8epr01qlsedscn00", nil)
	if err != nil {
		log.Panicf("dial failed %v", err)
	}
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"type":"subscribe","symbol":"BINANCE:BTCUSDT"}`)); err != nil {
		log.Panicf("write failed %v", err)
	}
	var resp Resp
	go func() {
		for {
			time.Sleep(time.Second * 10)
			_, data, err := c.ReadMessage()
			if err != nil {
				log.Printf("read: %v", err)
			}
			if err := json.Unmarshal(data, &resp); err != nil {
				log.Printf("unmarshal: %v", err)
				continue
			}
			d := resp.Data
			if len(resp.Data) > 1 {
				r := d[len(d)-1]
				f <- cast.ToFloat64(r.P)
			}
		}
	}()

	return f
}
