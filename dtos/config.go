package dtos

type Config struct {
	Database struct {
		Url string `json:"url" yaml:"url"`
	} `json:"database" yaml:"database"`
}
