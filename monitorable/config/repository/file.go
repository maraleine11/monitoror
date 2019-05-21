package repository

import (
	"os"

	"github.com/monitoror/monitoror/monitorable/config/models"
)

func (cr *configRepository) GetConfigFromPath(path string) (config *models.Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	config, err = GetConfig(file)
	return
}