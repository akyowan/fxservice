package adapter

import (
	"fxservice/service/logcenter/domain"
)

func DumpAdd(dump *domain.Dump) error {
	c := mgoPool.C("dumps")
	if err := c.Insert(dump); err != nil {
		return err
	}
	return nil
}
