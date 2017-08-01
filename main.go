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
<<<<<<< HEAD
			Text := update.Message.Text + "123"
=======
			Text := update.Message.Text + "  123"
>>>>>>> be0ed4bb57d0096856552db69f6a0e80b959cfd6

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Создаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			mm := SendMsg(msg)
			bot.Send(mm) //bot.SendMessage(msg)
		}

	}
}

func SendMsg(msg tgbotapi.MessageConfig) tgbotapi.Chattable {
	return msg
}
