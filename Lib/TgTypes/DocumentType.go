package TgTypes

type DocumentType struct {
	FileId       string        `json:"file_id"`
	FileUniqueId string        `json:"file_unique_id"`
	Thumb        PhotoSizeType `json:"thumb,omitempty"`
	FileName     string        `json:"file_name,omitempty"`
	MimeType     string        `json:"mime_type,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

type PhotoSizeType struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

type ChatPhotoType struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type StickerType struct {
	FileId       string           `json:"file_id"`
	FileUniqueId string           `json:"file_unique_id"`
	Width        int              `json:"width"`
	Height       int              `json:"height"`
	IsAnimated   bool             `json:"is_animated"`
	IsVideo      bool             `json:"is_video"`
	Thumb        PhotoSizeType    `json:"thumb,omitempty"`
	Emoji        string           `json:"emoji,omitempty"`
	SetName      string           `json:"set_name,omitempty"`
	MaskPosition MaskPositionType `json:"mask_position,omitempty"`
	FileSize     int              `json:"file_size,omitempty"`
}

type StickerSetType struct {
	Name          string        `json:"name"`
	Title         string        `json:"title"`
	IsAnimated    bool          `json:"is_animated"`
	IsVideo       bool          `json:"is_video"`
	ContainsMasks bool          `json:"contains_masks"`
	Stickers      []StickerType `json:"stickers"`
	Thumb         PhotoSizeType `json:"thumb,omitempty"`
}

type MaskPositionType struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type VideoType struct {
	FileId       string        `json:"file_id"`
	FileUniqueId string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Duration     int           `json:"duration"`
	Thumb        PhotoSizeType `json:"thumb,omitempty"`
	FileName     string        `json:"file_name,omitempty"`
	MimeType     string        `json:"mime_type,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

type VideoNoteType struct {
	FileId       string        `json:"file_id"`
	FileUniqueId string        `json:"file_unique_id"`
	Length       int           `json:"length"`
	Duration     int           `json:"duration"`
	Thumb        PhotoSizeType `json:"thumb,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

type AudioType struct {
	FileId       string        `json:"file_id"`
	FileUniqueId string        `json:"file_unique_id"`
	Duration     int           `json:"duration"`
	Performer    string        `json:"performer"`
	Title        string        `json:"title"`
	FileName     string        `json:"file_name"`
	MimeType     string        `json:"mime_type"`
	FileSize     int           `json:"file_size"`
	Thumb        PhotoSizeType `json:"thumb"`
}
