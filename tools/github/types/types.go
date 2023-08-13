package types

import "time"

type PR struct {
	Author              string    `json:"author"`
	CreatedAt           time.Time `json:"createdAt"`
	Id                  string    `json:"id"`
	Number              int       `json:"number"`
	Repository          string    `json:"repository"`
	RepositoryWithOwner string    `json:"repositoryWithOwner"`
	State               string    `json:"state"`
	Title               string    `json:"title"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type Repo struct {
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	FullName    string    `json:"fullName"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	OwnerID     string    `json:"ownerId"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}
