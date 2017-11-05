package adapter

import (
	"fxservice/service/logcenter/domain"
)

func FeedbackAdd(feedback *domain.FeedBack) error {
	c := mgoPool.C("feedbacks")
	if err := c.Insert(feedback); err != nil {
		return err
	}
	return nil
}
