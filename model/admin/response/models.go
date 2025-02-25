package response

type Model struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Assistant   string `json:"assistant"`
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Deleted     bool   `json:"deleted"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}
