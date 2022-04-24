package TgTypes

import (
	"fmt"
	"time"
)

var (
	colorReset = "\033[0m"
	bold       = "\u001b[1m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	//colorWhite  = "\033[37m"
)

func (u UserType) String() string {
	var res = colorRed
	if u.IsBot {
		res += fmt.Sprint(" BOT ")
	}
	res += fmt.Sprint(u.FirstName, " ")
	if u.LastName != "" {
		res += fmt.Sprint(u.LastName, " ")
	}
	if u.UserName != "" {
		res += fmt.Sprint(u.UserName, " ")
	}
	res += colorReset
	return res
}

func (u UpdateType) String() string {
	var res string

	if u.Message.MessageId != 0 {
		res += fmt.Sprint(u.Message)
	}
	if u.CallbackQuery.Id != "" {
		res += fmt.Sprint(u.CallbackQuery)
	}

	return res
}

func (c ChatType) String() string {
	var res string
	res += fmt.Sprint(c.Title, " ")
	return res
}

func (m MessageType) String() string {
	var res string
	// “private”, “group”, “supergroup” or “channel”
	res += bold
	if m.Chat.Type == "group" {
		res += colorPurple
	} else if m.Chat.Type == "supergroup" {
		res += colorCyan
	} else if m.Chat.Type == "private" {
		res += colorBlue
	} else {
		res += colorYellow
	}
	res += fmt.Sprint(m.MessageId, " ")
	res += fmt.Sprint("[", time.Unix(int64(m.Date), 0).String()[:19], "] ")
	res += fmt.Sprint(m.Chat)
	res += fmt.Sprint(m.From)
	if m.ForwardFromMessageId != 0 {
		res += fmt.Sprint("[forwarded from ", m.ForwardFromChat, "] ")
	}
	res += bold + colorBlue + "»»» " + colorGreen
	if m.ReplyToMessage != nil {
		res += fmt.Sprint("[reply to ", m.ReplyToMessage.MessageId, "] ")
	}

	if m.Animation != *new(AnimationType) {
		res += fmt.Sprint("[animation ", m.Animation.FileName, " size=", m.Animation.FileSize, "] ")
	} else if m.Audio != *new(AudioType) {
		res += fmt.Sprint("[audio ", m.Audio.FileName, " size=", m.Document.FileSize, "] ")
	} else if m.Document != *new(DocumentType) {
		res += fmt.Sprint("[document ", m.Document.FileName, " size=", m.Document.FileSize, "] ")
	} else if len(m.Photo) != 0 {
		res += fmt.Sprint("[photo] ")
	} else if m.Sticker != *new(StickerType) {
		res += fmt.Sprint("[sticker ", m.Sticker.Emoji, " size=", m.Sticker.FileSize, "] ")
	} else if m.Video != *new(VideoType) {
		res += fmt.Sprint("[video] ")
	} else if m.Voice != *new(VoiceType) {
		res += fmt.Sprint("[voice size=", m.Voice.FileSize, "] ")
	}

	if len(m.NewChatMembers) != 0 {
		res += fmt.Sprint("[new Members: ")
		for _, v := range m.NewChatMembers {
			res += v.String() + " "
		}
		res += "] "
	}

	if m.GroupChatCreated {
		res += "[group Chat Created] "
	}
	if m.SupergroupChatCreated {
		res += "[super Group Chat Created] "
	}
	if m.ChannelChatCreated {
		res += "[channel Chat created] "
	}
	if len(m.Entities) != 0 {
		for _, v := range m.Entities {
			res += v.String(m.Text)
		}
	}
	if len(m.CaptionEntities) != 0 {
		for _, v := range m.CaptionEntities {
			res += v.String(m.Caption)
		}
	}

	res += fmt.Sprint(m.Text, m.Caption) + colorReset
	return res
}

func (e MessageEntityType) String(text string) string {
	var res string
	if e.Type == "mention" {
		res += "[mention] "
	} else {
		res += "[" + e.Type + " " + text[e.Offset+1:e.Offset+e.Length] + "] "
	}
	return res
}

func (b CallbackQueryType) String() string {
	var res string

	res += bold
	res += colorYellow
	res += fmt.Sprint(b.Message.MessageId, " ")
	res += fmt.Sprint(b.Message.Chat)
	res += b.From.String() + " "
	res += bold + colorBlue + "»»» " + colorGreen
	res += b.Data
	res += colorReset

	return res
}
