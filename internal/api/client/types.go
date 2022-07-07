package client

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Agreement bool   `json:"agreement"`
	Locale    string `json:"locale"`
	Reason    string `json:"reason"`
}

type UpdateCredentialsRequest struct {
	Discoverable     string `json:"discoverable"`
	Bot              bool   `json:"bot"`
	DisplayName      string `json:"display_name"`
	Note             string `json:"note"`
	Avatar           string `json:"avatar"`
	Header           string `json:"header"`
	Locked           bool   `json:"locked"`
	SourcePrivacy    string `json:"source[privacy]"`
	SourceSensitive  bool   `json:"source[sensitive]"`
	SourceLanguage   string `json:"source[language]"`
	FieldsAttributes []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"fields[attributes]"`
}
