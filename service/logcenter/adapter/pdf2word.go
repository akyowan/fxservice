package adapter

import (
	"fxservice/service/logcenter/domain"
	"gopkg.in/mgo.v2/bson"
)

func GetLatestPDF2WordVersion() (*domain.Version, error) {
	var version domain.Version
	c := mgoPool.C("pdf2word_versions")
	if err := c.Find(nil).Sort("-version_seq").Limit(1).One(&version); err != nil {
		return nil, err
	}
	version.NeedUpdate = 1
	return &version, nil
}

func GetLatestForcePDF2WordVersion() (*domain.Version, error) {
	var version domain.Version
	c := mgoPool.C("pdf2word_versions")
	query := bson.M{
		"version_type": 1,
		"update_type": bson.M{
			"$in": []domain.UpdateTypeEnum{domain.UpdateTypePop, domain.UpdateTypeSilent},
		},
	}
	if err := c.Find(&query).Sort("-version_seq").Limit(1).One(&version); err != nil {
		return nil, err
	}
	return &version, nil
}
