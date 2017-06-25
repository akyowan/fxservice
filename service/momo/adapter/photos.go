package adapter

import (
	"fmt"
	"fxservice/domain"
)

func AddPhotos(photos [][]domain.Photos) error {
	db := dbPool.NewConn().Begin()
	beginID := 1000
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
