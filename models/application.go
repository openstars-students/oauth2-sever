package models

var (
	AppList map[string]*Application
)

type Application struct {
	ClientID               string
	ClientSecret           string
	ApplicationName        string
	HomepageURL            string
	ApplicationDescription string
	AuthorizeCallBackURL   string
	ImageURL               string
}