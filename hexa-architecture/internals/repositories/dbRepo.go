package repositories

import (
	"log"
	"os"

	"github.com/kaparouita/fiber_api/internals/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbRepo struct {
	Db *gorm.DB
}

/**  Init a new DbRepo and Connect to the database
 *  return the new DbRepo
 */
func ConnectDb() *DbRepo {
	//dsn := "host=localhost user=postgres password=kaparouita321 dbname=postgres port=5432 sslmode=disable"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	//TODO: update mysql

	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
		os.Exit(2)
	}

	log.Println("Successfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Mingrations")

	//Create the tables
	db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Order{})
	return &DbRepo{Db: db}
}

func CloseDb(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()
	return nil
}
