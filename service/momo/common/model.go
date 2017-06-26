package common

import (
	"math/rand"
	"time"
)

type PhoneModel struct {
	Model    string
	ModelNum []string
}

type Model struct {
	Model    string
	ModelNum string
}

var models []PhoneModel

func init() {
	models = []PhoneModel{
		{
			ModelNum: []string{"NNAC2", "NNAE2", "NNAF2", "NNAM2", "MNAJ2", "NNAP2", "MNAY2", "MN8L2", "MN8G2", "MN8V2", "MNGQ2", "MNGT2", "MNGX2", "MNGY2", "MNGW2", "MNC62", "MNH12", "MNH02", "MNC22"},
			Model:    "iPhone9,1",
		},
		{
			ModelNum: []string{"MN482", "MN4A2", "MN4C2", "MN4E2", "MN4K2", "MNRJ2", "MNRL2", "MNRM2", "MNQK2", "MNQH2", "MNFP2", "MNFR2", "MNFT2", "MNFQ2", "MNFV2", "MNFY2"},
			Model:    "iPhone9,2",
		},
		{
			ModelNum: []string{"NN912", "MN922", "MN942", "MN952", "MN962", "MN972", "MN992", "MN9D2", "MN9G2", "NN9H2", "MN9L2", "MN9N2", "MN9R2", "MN9U2", "MN9Y2", "MN8X2", "MN8Y2"},
			Model:    "iPhone9,3",
		},
		{
			ModelNum: []string{"MNQM2", "MNQT2", "MNQQ2", "MNQP2", "MN502", "MN532", "MN552", "MN562", "MN592", "MN5C2", "MN5D2", "MN5K2", "MN5M2", "MN5J2", "MN4M2", "MN4P2", "MN4U2", "MN4W2", "MN4Y2"},
			Model:    "iPhone9,4",
		},
	}
}

func GenRandModel() Model {
	var model Model
	rand.Seed(int64(time.Now().Nanosecond()))
	mlen := len(models)
	index := rand.Intn(mlen)

	model.Model = models[index].Model
	model.ModelNum = models[index].ModelNum[rand.Intn(len(models[index].ModelNum))]
	return model

}
