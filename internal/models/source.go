package models

type Source struct {
	ID                 uint64
	Note               string `json:"note"`
	Privacy            string `json:"privacy"`
	Sensitive          bool   `json:"sensitive"`
	Language           string `json:"language"`
	FollowRequestCount int    `json:"follow_request_count"`

	//Fields             []Field `json:"fields"`
}
