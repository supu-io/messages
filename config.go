package messages

// Github config
type Github struct {
	Token  string `json:"token"`
	Labels []string
}

// Config internal to be shared between services
type Config struct {
	Github *Github `json:"github, omitempty"`
}

// GetIssuesList request message to get an issue list
type GetIssuesList struct {
	Status string `json:"status"`
	Org    string `json:"org"`
	Repo   string `json:"repo,omitempty"`
	Config `json:"config"`
}

// IssuesDetails request message to get an issue
type GetIssue struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Config `json:"config"`
}

// CreateIssue request message to create an issue
type CreateIssue struct {
	Issue  *Issue `json:issue`
	Config `json:"config"`
}
