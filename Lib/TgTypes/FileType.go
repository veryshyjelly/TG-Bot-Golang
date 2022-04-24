package TgTypes

type FileType struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}
