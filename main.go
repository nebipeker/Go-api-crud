package main

import (
	"fmt"
	"log"
	"simple-res-api/controllers"
	"simple-res-api/models"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := readConfig()

	db := initDB(config)
	migrateDB(db)

	e := echo.New()

	controller := &controllers.Controller{DB: db}

	// Company routes
	e.GET("/api/companies", controller.GetCompanies)
	e.GET("/api/companies/:company_id", controller.GetCompanyById)
	e.POST("/api/companies", controller.CreateCompany)
	e.PUT("/api/companies/:company_id", controller.UpdateCompany)
	e.DELETE("/api/companies/:company_id", controller.DeleteCompany)

	// Job listings
	e.GET("/api/jobs", controller.GetJobListings)
	e.GET("/api/jobs/:job_id", controller.GetJobListingById)
	e.POST("/api/jobs", controller.CreateJobListing)
	e.PUT("/api/jobs/:job_id", controller.UpdateJobListing)
	e.DELETE("/api/jobs/:job_id", controller.DeleteJobListing)

	// Candidate profiles //DONE
	e.GET("/api/candidates", controller.GetCandidates)
	e.GET("/api/candidates/:candidate_id", controller.GetCandidateById)
	e.POST("/api/candidates", controller.CreateCandidate)
	e.PUT("/api/candidates/:candidate_id", controller.UpdateCandidate)
	e.DELETE("/api/candidates/:candidate_id", controller.DeleteCandidate)

	e.Start(":8080")
}

func readConfig() Config {
	viper.SetConfigName("docker-compose")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var config Config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Error unmarshaling config: %s \n", err))
	}
	return config
}

func initDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Database.Username, config.Database.Password,
		config.Database.DBName, config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Company{}, &models.JobListing{}, &models.Candidate{})
}

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}
