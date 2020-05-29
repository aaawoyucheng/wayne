package health

import (
	"github.com/aaawoyucheng/wayne/src/backend/models"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	_, err := models.Ormer().
		QueryTable(new(models.Cluster)).
		Count()
	if err != nil {
		return err
	}
	return nil
}
