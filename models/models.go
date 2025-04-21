package models

type Users struct {
	Data []User `json:"data"`
}

type User struct {
	GID          string      `json:"gid"`
	ResourceType string      `json:"resource_type"`
	Name         string      `json:"name"`
	Email        string      `json:"email, omitempty"`
	Photo        string      `json:"photo, omitempty"`
	Workspaces   []Workspace `json:"workspaces"`
}

type Workspaces struct {
	Data     []Workspace `json:"data"`
	NextPage Page        `json:"next_page"`
}

type Workspace struct {
	GID          string `json:"gid"`
	ResourceType string `json:"resource_type"`
	Name         string `json:"name"`
}

type Page struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}
