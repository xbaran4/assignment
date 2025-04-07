package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type User struct {
	ID          uint
	ExternalId  string
	Name        string
	Email       string
	DateOfBirth time.Time
}

type UserRepository interface {
	CreateUser(User) (User, error)
	GetUser(uint) (User, error)
}

type userRepository struct {
	*gorm.DB
}

func InitUserRepository(dbFilename string) UserRepository {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{IgnoreRecordNotFoundError: true})}) // Ignore ErrRecordNotFound error for logger

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	return userRepository{db}
}

func (r userRepository) CreateUser(user User) (User, error) {
	result := r.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r userRepository) GetUser(id uint) (User, error) {
	var user User
	result := r.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
