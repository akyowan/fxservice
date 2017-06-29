package adapter

import (
	"fmt"
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

func AddPhotos(photos [][]domain.Photos) error {
	db := dbPool.NewConn().Begin()
	beginID := 1120
	for i := range photos {
		beginID = beginID + 1
		photoGroup := photos[i]
		photoID := fmt.Sprintf("%d", beginID)
		for j := range photoGroup {
			photo := photoGroup[j]
			photo.PhotosID = photoID
			if err := db.Create(&photo).Error; err != nil {
				db.Rollback()
				return err
			}
		}
	}
	db.Commit()
	return nil
}

func GetPhotos(photosID string) ([]domain.Photos, error) {
	db := dbPool.NewConn()
	var photos []domain.Photos
	dbResult := db.Where("photos_id = ?", photosID).Order("seq").Find(&photos)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return photos, nil
}

func GetAvatar(photosID string) (string, error) {
	db := dbPool.NewConn()
	var photo domain.Photos
	dbResult := db.Where("photos_id = ?", photosID).Order("seq").First(&photo)
	if dbResult.RecordNotFound() {
		return "", errors.NotFound
	}
	if dbResult.Error != nil {
		return "", dbResult.Error
	}
	return photo.URL, nil
}

func GetRandomPhotosID() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(230) + 1000
	return fmt.Sprintf("%d", id)
}
