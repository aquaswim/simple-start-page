package entities

type Setting struct {
	SiteTitle   string `json:"siteTitle"`
	About       string `json:"siteAbout"`
	HideSetting bool   `json:"hideSetting"`
}
