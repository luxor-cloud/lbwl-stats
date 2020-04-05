package flash

import (
	"time"
)

type GlobalGameData struct {
	UUID              string `gorm:"primarykey;type:varchar(36)"`
	Wins              uint32 `gorm:"type:int unsigned"`
	Deaths            uint32 `gorm:"type:int unsigned"`
	CheckpointRecords uint32 `gorm:"type:int unsigned"`
	GamesPlayed       uint32 `gorm:"type:int unsigned"`
	InstantDeaths     uint32 `gorm:"type:int unsigned"`
}

type MapRecords struct {
	UUID           string              `￿￿gorm:"primarykey;type:varchar(36)"`
	Map            string              `gorm:"primarykey;type:varchar(100)"`
	RecordTime     uint64              `gorm:"type:bigint unsigned"`
	AccomplishedAt time.Time           `gorm:"type:datetime"`
	Checkpoints    []CheckpointRecords `gorm:"foreignkey:uuid;foreignkey:map"`
}

type CheckpointRecords struct {
	UUID           string    `gorm:"primarykey;type:varchar(36)"`
	Map            string    `gorm:"primarykey;type:varchar(100)"`
	Checkpoint     byte      `gorm:"type:tinyint unsigned"`
	RecordTime     uint64    `gorm:"type:bigint unsigned"`
	AccomplishedAt time.Time `gorm:"type:datetime"`
}

type Player struct {
	GlobalGameData GlobalGameData `gorm:"foreignkey:uuid"`
	MapRecords     []MapRecords   `gorm:"foreignkey:uuid;foreignkey:map"`
}
