package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read(filename string) ([]byte, error) {
	return nil, nil
}

func (db *CloudDb) Write(content, filename string) {

}
