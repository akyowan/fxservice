package adapter

import (
	"fmt"
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

const MAX_PHOTO_GROUP_RANDOM = 10000000

func AddPhotos(photoGroups [][]domain.Photo) error {
	db := dbPool.NewConn().Begin()
	beginID := time.Now().Unix()
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range photoGroups {
		beginID += 1
		photos := photoGroups[i]
		photosID := fmt.Sprintf("%x", beginID)
		random := rand.Intn(MAX_PHOTO_GROUP_RANDOM)
		photoGroup := domain.PhotoGroup{
			PhotosID: photosID,
			Status:   domain.PhotosStatusFree,
			Random:   random,
		}
		if err := db.Create(&photoGroup).Error; err != nil {
			db.Rollback()
			return err
		}
		for j := range photos {
			photo := photos[j]
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
	if err := db.Where("status = ?", domain.PhotosStatusFree).Order("random").First(&photoGroup).Error; err != nil {
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
