package api

import "github.com/PedroBett-dev/Golang-Api/src/config"

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	return nil
}
