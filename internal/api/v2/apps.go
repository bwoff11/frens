package v2

type CreateAppRequestBody struct {
	ClientName   string   `json:"client_name" form:"client_name" validate:"required,max=64"` // A name for your application
	RedirectURIs []string `json:"redirect_uris" form:"redirect_uris validate:"required"`     // Where the user should be redirected after authorization. To display the authorization code to the user instead of redirecting to a web page, use urn:ietf:wg:oauth:2.0:oob in this parameter.
	Scopes       []string `json:"scopes" form:"scopes"`                                      // Space separated list of scopes. If none is provided, defaults to read.
	Website      string   `json:"website" form:"website" validate:"max=64"`                  // A URL to the homepage of your app
}
