package config

import (
	"github.com/bugsnag/bugsnag-go/v2"
)

func BugsnagClient() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          Config.Bugsnag.ApiKey,
		ReleaseStage:    Config.Bugsnag.ReleaseStage,
		ProjectPackages: []string{"main", "api-gateway/cmd"},
	})
}
