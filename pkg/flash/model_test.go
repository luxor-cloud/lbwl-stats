package flash

import (
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type GlobalGameData struct {
	UUID              string       `gorm:"primarykey;type:varchar(36)"`
	Wins              uint32       `gorm:"type:int unsigned"`
	Deaths            uint32       `gorm:"type:int unsigned"`
	Checkpoints       uint32       `gorm:"type:int unsigned"`
	GamesPlayed       uint32       `gorm:"type:int unsigned"`
	InstantDeaths     uint32       `gorm:"type:int unsigned"`
	MapRecords        []MapRecords `gorm:"foreignkey:uuid;association_foreignkey:uuid"`
}

type MapRecords struct {
	UUID           string              `￿￿gorm:"primarykey;type:varchar(36)"`
	Map            string              `gorm:"primarykey;type:varchar(100)"`
	RecordTime     uint64              `gorm:"type:bigint unsigned"`
	AccomplishedAt time.Time           `gorm:"type:datetime"`
	Checkpoints    []CheckpointRecords `gorm:"foreignkey:uuid;foreignkey:map;association_foreignkey:uuid;association_foreignkey:map"`
}

type CheckpointRecords struct {
	UUID           string    `gorm:"primarykey;type:varchar(36)"`
	Map            string    `gorm:"primarykey;type:varchar(100)"`
	Checkpoint     byte      `gorm:"type:tinyint unsigned"`
	RecordTime     uint64    `gorm:"type:bigint unsigned"`
	AccomplishedAt time.Time `gorm:"type:datetime"`
}

func Test_TestTheShit(t *testing.T) {
	db, err := gorm.Open("mysql", "root:secret@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	l := GlobalGameData{}


	err = db.Preload("MapRecords.Checkpoints").Find(&l, "uuid = ?", "92de217b-8b2b-403b-86a5-fe26fa3a9b5f").Error//db.Preload("MapRecords").Find(&l).Error
	if err != nil {
		log.Fatal(err)
	}

	//for _, p := range l {

		log.Printf("UUID: %v\n", l.UUID)
		log.Printf("GamesPlayed: %v\n", l.GamesPlayed)
		log.Printf("Wins: %v\n", l.Wins)
		log.Printf("Deaths: %v\n", l.Deaths)
		log.Printf("Instant: %v\n", l.InstantDeaths)
		log.Printf("Check: %v\n", l.Checkpoints)


		for _, mr := range l.MapRecords {
			log.Printf("  Map: %v\n", mr.Map)
			log.Printf("  Accomp: %v\n", mr.AccomplishedAt)
			log.Printf("  Rec: %v\n", mr.RecordTime)

			for _, cr := range mr.Checkpoints {
				log.Printf("    Accomp: %v\n", cr.AccomplishedAt)
				log.Printf("    Rec: %v\n", cr.RecordTime)
			}
		}

	//}
}
