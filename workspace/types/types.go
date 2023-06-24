package types

import "time"

type Config struct {
	Frequency  time.Duration `json:"frequency"`
	OutputFile string        `json:"output_file"`
}

type Data struct {
	CreatedAt string `json:"created_at"`
	IconURL   string `json:"icon_url"`
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
	Value     string `json:"value"`
}
