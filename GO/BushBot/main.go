package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Функция входа
func main() {
	botToken := "1613023350:AAEDY59XvcXlVpgJqqs6aDpQukGadVjMftU"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	//offset := 0
	for {
		updates, err := getUpdates(botUrl)
		if err != nil {
			log.Println("Что-то не так: ", err.Error())
		}
		// for _, update := range updates {
		// 	err = respond(botUrl, update)
		// 	offset = update.UpdateId + 1
		// }
		fmt.Println(updates)
	}
}

// func handleRequest(){
// запрос обновлений
func getUpdates(botUrl string) (Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// var restResponse RestResponse
	// err = json.Unmarshal(body, &restResponse)
	// if err != nil {
	// 	return nil, err
	// }
	// return restResponse.Result, nil
}

// ответ на обновления
func respond() {
	// var botMessage BotMessage
	// botMessage.ChatId = update.Message.Chat.ChatId
	// botMessage.Text = update.Message.Text
	// buf, err := json.Marshal(botMessage)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	// if err != nil {
	// 	panic(err)
	// }
	// return nil
}
