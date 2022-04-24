package TgTypes

type PollType struct {
	Id                    string              `json:"id"`
	Question              string              `json:"question"`
	Options               []PollOptionType    `json:"options"`
	TotalVoterCount       int                 `json:"total_voter_count"`
	IsClosed              bool                `json:"is_closed"`
	IsAnonymous           bool                `json:"is_anonymous"`
	Type                  string              `json:"type"` // Regular or Quiz
	AllowsMultipleAnswers bool                `json:"allows_multiple_answers"`
	CorrectOptionId       int                 `json:"correct_option_id,omitempty"`
	Explanation           string              `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntityType `json:"explanation_entities,omitempty"`
	OpenPeriod            int                 `json:"open_period,omitempty"`
	CloseDate             int                 `json:"close_date,omitempty"`
}

type PollOptionType struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PollAnswerType struct {
	PollId    string   `json:"poll_id"`
	User      UserType `json:"user"`
	OptionIds []int    `json:"option_ids"`
}
