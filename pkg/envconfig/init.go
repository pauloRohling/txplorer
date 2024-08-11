package envconfig

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/pauloRohling/txplorer/pkg/yml"
	"os"
)

const ProfileEnv = "PROFILE"

func Init(v any) (Profile, error) {
	if err := yml.Read("./application.yml", v); err != nil {
		return "", err
	}

	profile := NewProfile(os.Getenv(ProfileEnv))
	profileConfigurationFile := fmt.Sprintf("./application-%s.yml", profile)

	if err := yml.Read(profileConfigurationFile, v); err != nil {
		return "", err
	}

	if err := envconfig.Process("", v); err != nil {
		return "", err
	}

	return NewProfile(os.Getenv(ProfileEnv)), nil
}
