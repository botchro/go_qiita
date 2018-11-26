package main

// Tags an array of tag
type Tags []struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}
