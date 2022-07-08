package models

type Source struct {
	Note               string  `json:"note"`
	Fields             []Field `json:"fields"`
	Privacy            string  `json:"privacy"`
	Sensitive          bool    `json:"sensitive"`
	Language           string  `json:"language"`
	FollowRequestCount int     `json:"follow_request_count"`
}
