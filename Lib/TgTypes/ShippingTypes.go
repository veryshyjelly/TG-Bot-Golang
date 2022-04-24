package TgTypes

type ShippingQueryType struct {
	Id              string              `json:"id"`
	From            UserType            `json:"from"`
	InvoicePayload  string              `json:"invoice_payload"`
	ShippingAddress ShippingAddressType `json:"shipping_address"`
}

type PreCheckoutQueryType struct {
	Id               string        `json:"id"`
	From             UserType      `json:"from"`
	Currency         string        `json:"currency"`
	TotalAmount      int           `json:"total_amount"`
	InvoicePayload   string        `json:"invoice_payload"`
	ShippingOptionId string        `json:"shipping_option_id,omitempty"`
	OrderInfo        OrderInfoType `json:"order_info"`
}

type ShippingAddressType struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}

type OrderInfoType struct {
	Name            string              `json:"name,omitempty"`
	PhoneNumber     string              `json:"phone_number,omitempty"`
	Email           string              `json:"email,omitempty"`
	ShippingAddress ShippingAddressType `json:"shipping_address,omitempty"`
}
