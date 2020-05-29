package models_test

import (
	"testing"
	"time"

	"github.com/aaawoyucheng/wayne/src/backend/models"
)

func TestGetDeployCountByDay(t *testing.T) {
	result, err := models.PublishHistoryModel.GetDeployCountByDay(time.Now().Add(-24*30*time.Hour), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
