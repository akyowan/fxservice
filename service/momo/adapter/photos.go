package adapter

import (
	"fmt"
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

func AddPhotos(photos [][]domain.Photo) error {
	db := dbPool.NewConn().Begin()
	beginID := time.Now().Unix()
	for i := range photos {
		beginID += 1
		photoGroup := photos[i]
		photosID := fmt.Sprintf("%x", beginID)
		photos := domain.PhotoGroup{
			PhotosID: photosID,
			Status:   domain.PhotosStatusFree,
		}
		if err := db.Create(&photos).Error; err != nil {
			db.Rollback()
			return err
		}
		for j := range photoGroup {
			photo := photoGroup[j]
			photo.PhotosID = photosID
			if err := db.Create(&photo).Error; err != nil {
				db.Rollback()
				return err
			}
		}
	}
	db.Commit()
	return nil
}

func GetPhotos(photosID string) ([]domain.Photo, error) {
	db := dbPool.NewConn()
	var photos []domain.Photo
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
	var photo domain.Photo
	dbResult := db.Where("photos_id = ?", photosID).Order("seq").First(&photo)
	if dbResult.RecordNotFound() {
		return "", errors.NotFound
	}
	if dbResult.Error != nil {
		return "", dbResult.Error
	}
	return photo.URL, nil
}

func GetFreeAvatar() (*domain.Photo, error) {
	db := dbPool.NewConn().Begin()
	var photoGroup domain.PhotoGroup
	if err := db.Where("status = ?", domain.PhotosStatusFree).First(&photoGroup).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	photoGroup.Status = domain.PhotosStatusUsed
	if err := db.Save(&photoGroup).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	var photo domain.Photo
	if err := db.Where("photos_id = ?", photoGroup.PhotosID).Order("seq").First(&photo).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return &photo, nil
}

func GetRandomPhotosID() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(230) + 1000
	return fmt.Sprintf("%d", id)
}
