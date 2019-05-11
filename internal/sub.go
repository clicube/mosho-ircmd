package internal

import (
	"os/exec"
	"log"

	cmdsub "mosho-cmdsub/pkg"
)

func Sub() error{
	sub, err := cmdsub.New()
	if err != nil {
		return err
	}

	for {
		data, err := sub.Next()
		if err != nil {
			return err
		}

		path := "./onCmd.sh"
		out, err := handleScript(path, data)
		if err != nil {
			return err
		}
		log.Printf("%s: %s", path, out)
	}
}

func handleScript(path string, data map[string]interface{}) (string, error) {
	out, err := exec.Command(path).Output()
	return string(out), err
}
