package balance_model

import (
	"github.com/jinzhu/gorm"
	"github.com/seannguyen/coin-tracker/internal/models/snapshot_model"
)

type Balance struct {
	gorm.Model
	SnapshotID int
	Snapshot   snapshot_model.Snapshot
	Currency   string `gorm:"size:10;index"`
	Amount     float64
}
