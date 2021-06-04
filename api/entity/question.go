package entity

type Response struct {
	ResponseCode int       `json:"response_code"`
	Results      []Results `json:"results"`
}

type Results struct {
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	IncorrectAnswer []string `json:"incorrect_answers"`
}
