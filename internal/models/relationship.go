package models

type Relationship struct {
	ID              uint64 `json:"id"`
	SourceAccountID uint64 `json:"source_account_id"`
	SourceAccount   string `json:"source_account"`
	TargetAccountID uint64 `json:"target_account_id"`
	TargetAccount   string `json:"target_account"`
	Following       bool   `json:"following"`
	Blocking        bool   `json:"blocking"`
	Muting          bool   `json:"muting"`
}
