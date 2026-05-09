package global

import (
	"os"
	"path/filepath"
	"strings"
)

func GetConfigPath() (string, error) {
	if GlobalConfigPath != "" {
		return ExpandPath(GlobalConfigPath), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".rex", "config.yaml"), nil
}

func ExpandPath(path string) string {

	if path == "" {
		return path
	}

	// note: ~/xxx
	if path == "~" || strings.HasPrefix(path, "~/") {

		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		if path == "~" {
			return home
		}

		return filepath.Join(
			home,
			path[2:],
		)
	}

	// note: env expand
	// Linux/macOS:
	//   $HOME/test.yaml
	//
	// Windows:
	//   %USERPROFILE%\test.yaml
	path = os.ExpandEnv(path)

	return path
}
