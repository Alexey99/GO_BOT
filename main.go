package main
import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
)

func main() {
  // ������������ � ���� � ������� ������
  bot, err := tgbotapi.NewBotAPI("�����")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// �������������� �����, ���� ����� ��������� ���������� �� API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	err = bot.UpdatesChan(ucfg)
	// ������ ���������� �� ������
	for {
		select {
		case update := <-bot.Updates:
			// ������������, ������� ������� ����
			UserName := update.Message.From.UserName

			// ID ����/�������.
			// ����� ���� ��������������� ��� ���� � �������������
			// (����� �� ����� UserID) ��� � ���������� ����/������
			ChatID := update.Message.Chat.ID

			// ����� ���������
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// ������� ������������ ��� �� ����������
			reply := Text
			// �������� ���������
			msg := tgbotapi.NewMessage(ChatID, reply)
			// � ���������� ���
			bot.SendMessage(msg)
		}

	}
}