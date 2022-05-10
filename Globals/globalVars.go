package Globals

import "Telegram-Bot/Lib/TgTypes"

var (
	AudioLinks = make(map[int64]string)
	VideoLinks = make(map[int64]string)

	AudioButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Audio 🎵",
		CallbackData: "ytAudio",
	}
	VideoButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Video 📽️",
		CallbackData: "ytVideo",
	}
	ExitButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Exit",
		CallbackData: "deleteMessage",
	}
	StickerButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Stickers",
		CallbackData: "stickerMenu",
	}
	FilterButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Filter",
		CallbackData: "filterMenu",
	}
	BackButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Back",
		CallbackData: "GoBack",
	}

	VioAPIData       = make(map[string]VioFeatureType)
	PhotoFilterQueue = make(map[int64]string)
)
