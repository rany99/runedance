package minio

import "runedance/common/config"

type CosVideo struct {
	VideoBucket string
	CoverBucket string
	SecretID    string
	SecretKey   string
}

var cosVideo CosVideo

func Init() {
	cosVideo = CosVideo{
		VideoBucket: config.COS.VideoBucket,
		CoverBucket: config.COS.CoverBucket,
		SecretID:    config.COS.SecretID,
		SecretKey:   config.COS.SecretKey,
	}
}
