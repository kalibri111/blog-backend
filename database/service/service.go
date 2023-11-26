package service

import (
	"backend/database/model"
	"gorm.io/gorm"
	"time"
)

func GetAllArticles(db *gorm.DB) ([]model.Article, error) {
	var articles []model.Article
	err := db.Model(&model.Article{}).Find(&articles).Error
	return articles, err
}

func GetAllPhotos(db *gorm.DB) ([]model.Photo, error) {
	var photos []model.Photo
	err := db.Model(&model.Photo{}).Find(&photos).Error
	return photos, err
}

func AddArticle(db *gorm.DB, article *model.Article) error {
	return db.Model(&model.Article{}).Save(article).Error
}

func AddPhoto(db *gorm.DB, photo *model.Photo) error {
	return db.Model(&model.Photo{}).Save(photo).Error
}

func AddUser(db *gorm.DB, user *model.User) error {
	return db.Model(&model.User{}).Save(user).Error
}

func GetArticleByHeader(db *gorm.DB, header string) (model.Article, error) {
	var article model.Article
	err := db.Model(&model.Article{}).Where(&model.Article{Header: header}).Find(&article).Error
	return article, err
}

func GetPhotosByDate(db *gorm.DB, datetime string) ([]model.Photo, error) {
	t, terr := time.Parse(time.RFC3339, datetime)

	if terr != nil {
		return nil, terr
	}

	var photos []model.Photo
	derr := db.Model(&model.Photo{}).Where(&model.Photo{Date: t}).Find(&photos).Error
	return photos, derr
}

func GetUserByFirstNameByLastName(db *gorm.DB, firstName string, lastName string) ([]model.User, error) {
	var users []model.User
	err := db.Model(&model.User{}).Where(&model.User{FirstName: firstName, LastName: lastName}).Find(&users).Error
	return users, err
}

func GetUserByCreds(db *gorm.DB, login string, password string) (model.User, error) {
	var user model.User
	err := db.Model(&model.User{}).Where(&model.User{Login: login, Password: password}).Find(&user).Error
	return user, err
}
