package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("441163061:AAGE7czgn1XHFvU8YN-PbyduhWlMbE9AGxQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	ch, err := bot.GetUpdatesChan(ucfg) //UpdatesChan(ucfg)

	msg := tgbotapi.NewMessage(402503192, "Hello")
	mm := SendMsg(msg)
	bot.Send(mm)

	// читаем обновления из канала
	for {
		select {
		case update := <-ch:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text + "123"

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Создаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)

			//tgbotapi.ReplyKeyboardMarkup.ResizeKeyboard
			//tgbotapi.KeyboardButton.Text kk := "123"
			//msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{Keyboard: kk}
			//tgbotapi.KeyboardButton.
			//var k1 = tgbotapi.KeyboardButton{Text: "123"}
			//var k2 = tgbotapi.KeyboardButton{Text: "123"}

			//var test string
			//test = "TEST"

			var ya string
			ya = "http://yandex.ru"

			var k1 = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Contact_1", RequestContact: true}}
			var k2 = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Location_2", RequestLocation: true}}

			//var k3 = []tgbotapi.InlineKeyboardButton{tgbotapi.InlineKeyboardButton{Text: "Push", CallbackData: &test}}
			var k4 = []tgbotapi.InlineKeyboardButton{tgbotapi.InlineKeyboardButton{Text: "Yandex", URL: &ya}}

			kk := [][]tgbotapi.KeyboardButton{}

			kk = append(kk, k1)
			kk = append(kk, k2)

			kk_ := [][]tgbotapi.InlineKeyboardButton{}

			//kk_ = append(kk_, k3)
			kk_ = append(kk_, k4)

			var keyboard = tgbotapi.ReplyKeyboardMarkup{Keyboard: kk, ResizeKeyboard: true} //[][]string{[]string{"Hi, Bot!"}}}

			var inkeyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: kk_}

			msg.ReplyMarkup = keyboard
			msg.ReplyMarkup = inkeyboard

			// и отправляем его
			mm := SendMsg(msg)
			bot.Send(mm) //bot.SendMessage(msg)
		}

	}
}

func SendMsg(msg tgbotapi.MessageConfig) tgbotapi.Chattable {
	return msg
}
