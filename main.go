package main

import (
	"github.com/spf13/viper"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

var db *gorm.DB

type Snapshot struct {
	gorm.Model
}

type Balance struct {
	gorm.Model
	SnapshotID int
	Snapshot Snapshot
	Currency string `gorm:"size:10;index"`
	Amount float64
}

func main() {
	initialize()
	defer destroy()

	//balances, err := bittrex.GetBalances()
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(balances)

	db.AutoMigrate(&Snapshot{}, &Balance{})

	//db.Create(&Snapshot{gorm.Model{ CreatedAt: time.Now(), UpdatedAt: time.Now() } })

	var snap Snapshot
	db.Take(&snap)
	fmt.Println(snap.CreatedAt)
}

func initialize() {
	initConfigs()
	initDatabase()
}

func initConfigs() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}

func initDatabase() {
	database, err := gorm.Open("postgres", viper.Get("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic(err)
	}
	log.Println("successfully connected to db")
	db = database
}

func destroy() {
	db.Close()
}