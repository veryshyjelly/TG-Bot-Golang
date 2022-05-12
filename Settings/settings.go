package Settings

import (
	"strconv"
	"strings"
)

const ApiToken = "5284202120:AAFFwt36cQNllO9UWcx5vy48RnG1UNMuyGs"
const BotId = "@Khonshu22Bot"
const BotName = "Khonshu"
const OwnerId = 1653921867
const VioKey = "73dc-6d6c-9e3e"
const VioKey2 = "4abf-1b93-50bd"
const Delay = 10
const BaseUrl = "https://api.telegram.org/bot" + ApiToken
const MenuText = "Menu\n\nHey! My name is <b>Khonshu</b> \n I'm here to help you manage your group.\nI have lots of features like making sticker from image, and predetermined replies on certain keywords."

var BotUserId, _ = strconv.Atoi(strings.Split(ApiToken, ":")[0])
