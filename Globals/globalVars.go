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
	YtLinkButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Get Links",
		CallbackData: "ytLinks",
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
	PhotoButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Photo",
		CallbackData: "photoMenu",
	}
	YoutubeButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Youtube",
		CallbackData: "youtubeMenu",
	}
	BackButton = TgTypes.InlineKeyboardButtonType{
		Text:         "Back",
		CallbackData: "GoBack",
	}

	VioAPIData       = make(map[string]VioFeatureType)
	PhotoFilterQueue = make(map[int64]string)
)
