package requests

type Country struct {
	Country             string `json:"Country"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
