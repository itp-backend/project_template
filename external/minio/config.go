package minio

type ClientConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Region     string
	BucketName string
}
