package TgTypes

type VoiceType struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type VoiceChatScheduledType struct {
	StartDate int `json:"start_date"`
}

type VoiceChatStartedType struct {
}

type VoiceChatParticipantsInvitedType struct {
	Users []UserType `json:"users,omitempty"`
}

type VoiceChatEndedType struct {
	Duration int `json:"duration"`
}
