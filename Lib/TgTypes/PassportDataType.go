package TgTypes

type PassportDataType struct {
	Data        []EncryptedPassportElementType `json:"data"`
	Credentials EncryptedCredentialsType       `json:"credentials"`
}

type PassportFileType struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElementType struct {
	Type        string             `json:"type"`
	Data        string             `json:"data,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty"`
	Email       string             `json:"email,omitempty,omitempty"`
	Files       []PassportFileType `json:"files,omitempty"`
	FrontSize   PassportFileType   `json:"front_size,omitempty"`
	ReverseSide PassportFileType   `json:"reverse_side,omitempty"`
	Selfie      PassportFileType   `json:"selfie,omitempty"`
	Translation []PassportFileType `json:"translation,omitempty"`
	Hash        string             `json:"hash"`
}

type EncryptedCredentialsType struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}
