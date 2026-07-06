package model

type CommonModel struct {
	RuntimePath string
}

type APPModel struct {
	LogPath           string
	LogSaveName       string
	LogFileExt        string
	AppStorePath      string
	AppsPath          string
	AutoUpdateEnabled bool
}

type ServerModel struct {
	AppStoreList []string `ini:"appstore,,allowshadow"`
}

type GlobalModel struct {
	OpenAIAPIKey string
}

type CasaOSGlobalVariables struct {
	AppChange bool
}
