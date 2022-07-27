package database

import (
	"TTU_CORE_ERP_API/models"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var err error

func loadCredentials() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"username": os.Getenv("DB_USERNAME"),
		"password": os.Getenv("DB_PASSWORD"),
		"host":     os.Getenv("DB_HOST"),
		"db_name":  os.Getenv("DB_NAME"),
		"port":     os.Getenv("DB_PORT"),
	}
}

func Init() *gorm.DB {
	var _ error
	credentials := loadCredentials()
	//Connect to db using GORM
	host := credentials["host"]
	port := credentials["port"]
	user := credentials["username"]
	password := credentials["password"]
	dbName := credentials["db_name"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)
	dbInstance, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if dbErr != nil {
		errors.New("error connecting to data")
	}

	DB = dbInstance
	return dbInstance
}

func Migrate() error {
	DB.AutoMigrate(
		&models.User{}, &models.Level{}, &models.Address{},
	)

	err = DB.AutoMigrate(&models.Faculty{})
	err = DB.AutoMigrate(&models.Affiliation{})
	err = DB.AutoMigrate(&models.Department{})
	err = DB.AutoMigrate(&models.Programme{})
	err = DB.AutoMigrate(&models.GradingSystem{})

	err = DB.AutoMigrate(&models.StudentConstraint{})
	err = DB.AutoMigrate(&models.Hall{})
	err = DB.AutoMigrate(&models.Course{})
	err = DB.AutoMigrate(&models.MountedCourse{})
	err = DB.AutoMigrate(&models.Examination{})
	err = DB.AutoMigrate(&models.Religion{})
	err = DB.AutoMigrate(&models.Nationality{})
	err = DB.AutoMigrate(&models.Region{})
	err = DB.AutoMigrate(&models.District{})
	err = DB.AutoMigrate(&models.Hall{})
	err = DB.AutoMigrate(&models.Disability{})
	err = DB.AutoMigrate(&models.Student{})
	err = DB.AutoMigrate(&models.Subject{})
	err = DB.AutoMigrate(&models.Issue{})
	err = DB.AutoMigrate(&models.Bank{})
	err = DB.AutoMigrate(&models.Payment{})
	err = DB.AutoMigrate(&models.FormerSchool{})
	err = DB.AutoMigrate(&models.Classing{})
	err = DB.AutoMigrate(&models.GradingSystem{})
	err = DB.AutoMigrate(&models.MountedCourse{})
	err = DB.AutoMigrate(&models.AcademicRecord{})
	err = DB.AutoMigrate(&models.Clearance{})
	if err != nil {
		return err
	}
	return nil
}
