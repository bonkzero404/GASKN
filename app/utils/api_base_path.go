package utils

import (
	"github.com/bonkzero404/gaskn/config"
)

// ApiBasePath /*
func ApiBasePath() string {
	var str string

	if config.Config("API_WRAP") != "" {
		str = str + "/" + config.Config("API_WRAP")
	}

	if config.Config("API_VERSION") != "" {
		str = str + "/" + config.Config("API_VERSION")
	}
	// str = fmt.Sprintf("/%s/%s", config.Config("API_WRAP"), config.Config("API_VERSION"))
	return str
}

func ApiClientBasePath() string {
	var strSub string

	strSub = ApiBasePath()

	if config.Config("API_CLIENT") != "" {
		strSub = strSub + "/" + config.Config("API_CLIENT")
	}
	return strSub + "/:" + config.Config("API_CLIENT_PARAM")
}
