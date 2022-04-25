package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"strings"
)

func ParseMessage(message *TgTypes.MessageType) (int64, int64, string, string, string) {
	chatId, messageId := message.Chat.Id, message.MessageId
	textBody, command, joinedArgument := message.Text, "", ""

	if len(textBody) > 0 && isPrefix(textBody[0]) {
		command = strings.Split(textBody[1:], " ")[0]
	}

	if split := strings.Split(textBody, " "); len(split) > 1 {
		joinedArgument = strings.Join(split[1:], " ")
	}

	return chatId, messageId, textBody, command, joinedArgument
}

func isPrefix(r byte) bool {
	return r == '#' || r == '/' || r == '.'
}
