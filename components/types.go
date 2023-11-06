package components

type link struct {
	URL string `json:"url"`
}

type LinkMetadata struct {
	Description string `json:"description"`
	Address     link   `json:"address"`
}
