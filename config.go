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
	Issue  *Issue `json:"issue"`
	Config `json:"config"`
}

// CreateIssue request message to create an issue
type CreateIssue struct {
	Issue  *Issue `json:issue`
	Config `json:"config"`
}

// Setup request message to setup issue tracker
type Setup struct {
	Org    string   `json:"org"`
	Repo   string   `json:"repo"`
	States []string `json:"states"`
	Config `json:"config"`
}

// UpdateIssue request message to update an issue
type UpdateIssue struct {
	Issue    *Issue `json:issue`
	Status   string `json:"state"`
	Config   `json:"config"`
	Workflow `json:"workflow"`
}

// Workflow ...
type Workflow struct {
	Transitions []Transition `json:"transitions"`
}

// Transition json representation
type Transition struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Hooks []Hook `json:"hooks"`
}

type Hook struct {
	URL string `json:"url"`
}
