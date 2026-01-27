package dtos

type Config struct {
	Database struct {
		Url string `json:"url" yaml:"url"`
	} `json:"database" yaml:"database"`
	Google GoogleSheets `yaml:"google"`
}

type GoogleSheets struct {
	CredentialsEnv []byte `yaml:"credentials"`
	SheetIDEnv     string `yaml:"id"`
}
