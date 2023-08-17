package infra

import (
	"fmt"
	"log"

	domainBook "service-api/src/main/app/book/domain"
	domainGenre "service-api/src/main/app/genre/domain"
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"
	domainTransaction "service-api/src/main/app/transactions/domain"
	domainUser "service-api/src/main/app/users/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	env, err := LoadEnv("../..")
	if err != nil {
		log.Fatal(err.Error())
	}

	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", env.Host, env.Port, env.User, env.Dbname, env.Pass)
	dsn, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to connect to database")
	}

	return dsn
}

func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&domainBook.Book{},
		&domainGenre.Genre{},
		&domainUser.User{},
		&domainTransaction.Transaction{},
		&domainPaymentMethod.PaymentMethod{},
	)
}
