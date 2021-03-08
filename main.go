package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost", port, "で起動中...")

	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := `
	Hello World!
	
	Envs: 
	  PORT: %s
	`
	fmt.Fprintf(
		w,
		msg,
		os.Getenv("PORT"),
	)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	// BOTを初期化
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// リクエストからBOTのイベントを取得
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
	}

	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}

			// メッセージが位置情報の場合
			case *linebot.LocationMessage:
				sendRestInfo(bot, event)
			}
		}
	}
}

func sendRestInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := getRestInfo(lat, lng)

	_, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do()
	if err != nil {
		log.Print(err)
	}
}

// APIレスポンス
type response struct {
	Results results `json:"results"`
}

// APIレスポンスの内容
type results struct {
	Shop []shop `json:"shop"`
}

// レストラン一覧
type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func getRestInfo(lat string, lng string) string {
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s",
		apiKey,
		lat,
		lng,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	info := ""
	for _, shop := range data.Results.Shop {
		info += shop.Name + "\n" + shop.Address + "\n\n"
	}

	return info
}
