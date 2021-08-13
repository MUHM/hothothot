package baseTypes

type PasswordConfig struct {
	Secret  string `json:"secret"`
	Default string `json:"default"`
}
