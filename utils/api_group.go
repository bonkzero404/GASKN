package utils

import (
	"fmt"
	"github.com/bonkzero404/gaskn/config"
)

// SetupApiGroup /*
func SetupApiGroup() string {
	str := fmt.Sprintf("/%s/%s", config.Config("API_WRAP"), config.Config("API_VERSION"))
	return str
}

func SetupSubApiGroup() string {
	return SetupApiGroup() + "/" + config.Config("API_CLIENT") + "/:" + config.Config("API_CLIENT_PARAM")
}
