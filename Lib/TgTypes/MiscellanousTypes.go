package TgTypes

type BotCommandType struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type DiceType struct {
	Emoji string `json:"emoji"`
	Value string `json:"value"`
}

type LocationType struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type VenueType struct {
	Location        LocationType `json:"location"` // Venue Location
	Title           string       `json:"title"`
	Address         string       `json:"address"`
	FoursquareId    string       `json:"foursquare_id,omitempty"`
	FoursquareType  string       `json:"foursquare_type,omitempty"`
	GooglePlaceId   string       `json:"google_place_id,omitempty"`
	GooglePlaceType string       `json:"google_place_type"`
}

type ProximityAlertTriggeredType struct {
	Traveller UserType `json:"traveller"`
	Watcher   UserType `json:"watcher"`
	Distance  int      `json:"distance"`
}

type InlineKeyboardMarkupType struct {
	InlineKeyboard [][]InlineKeyboardButtonType `json:"inline_keyboard"`
}

type InlineKeyboardButtonType struct {
	Text                         string           `json:"text"`
	Url                          string           `json:"url"`
	LoginUrl                     LoginUrlType     `json:"login_url,omitempty"`
	CallbackData                 string           `json:"callback_data,omitempty"`
	SwitchInlineQuery            string           `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string           `json:"switch_inline_query_current_chat"`
	CallbackGame                 CallBackGameType `json:"callback_game,omitempty"`
	Pay                          bool             `json:"pay,omitempty"`
}

type LoginUrlType struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SuccessfulPaymentType struct {
	Currency                string        `json:"currency"`
	TotalAmount             int           `json:"total_amount"`
	InvoicePayload          string        `json:"invoice_payload"`
	ShippingOptionId        string        `json:"shipping_option_id,omitempty"`
	OrderInfo               OrderInfoType `json:"order_info,omitempty"`
	TelegramPaymentChargeId string        `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string        `json:"provider_payment_charge_id"`
}

type InlineQueryType struct {
	Id       string       `json:"id"`
	From     UserType     `json:"from"`
	Query    string       `json:"query"`
	Offset   string       `json:"offset"`
	ChatType string       `json:"chat_type,omitempty"`
	Location LocationType `json:"location,omitempty"`
}

type ChosenInlineResultType struct {
	ResultId        string       `json:"result_id"`
	From            UserType     `json:"from"`
	Location        LocationType `json:"location,omitempty"`
	InlineMessageId string       `json:"inline_message_id,omitempty"`
	Query           string       `json:"query"`
}

type ContactType struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type InvoiceType struct {
}
