package adapter

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func PDFLogInput(logs []map[string]interface{}, ip string) error {
	now := time.Now()
	c := mgoPool.C("pdf_report")
	for _, log := range logs {
		//loggers.Info.Println(log)
		log["client_ip"] = ip
		log["log_time"] = &now
		if err := c.Insert(bson.M(log)); err != nil {
			return err
		}
	}

	return nil
}
