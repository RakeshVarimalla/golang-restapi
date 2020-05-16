package github



type  GitHubErrorResponse struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Code     string `json:"code"`
		Field    string `json:"field"`
		Message  string `json:"message"`
	} `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
}