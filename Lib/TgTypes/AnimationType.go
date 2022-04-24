package TgTypes

type AnimationType struct {
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
