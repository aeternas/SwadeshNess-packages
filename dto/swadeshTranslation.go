package translation

type SwadeshTranslation struct {
	Results []GroupTranslation
	Credits string
}

type LanguageTranslation struct {
	Name        string
	Translation string
	Cached      bool `json:"isCached"`
}

type GroupTranslation struct {
	Name    string
	Results []LanguageTranslation
}
