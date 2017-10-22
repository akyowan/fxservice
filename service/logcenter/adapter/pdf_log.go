package adapter

import (
	//"fxlibraries/loggers"
	"gopkg.in/mgo.v2/bson"
)

func PDFLogInput(logs []map[string]interface{}) error {
	c := mgoPool.C("pdf_report")

	for _, log := range logs {
		//loggers.Info.Println(log)
		if err := c.Insert(bson.M(log)); err != nil {
			return err
		}
	}

	return nil
}
