package language

type Language struct {
	FullName string `json:"fullName"`
	Code     string
}

type LanguageGroup struct {
	Name      string
	Languages []Language
}
