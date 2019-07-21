package datastruct

//News ..
type News struct {
	ID      int64
	Author  string
	Body    string
	Created string
}

//NewsResponse data
type NewsResponse struct {
	ResponseCode string       `json:"responseCode"`
	ResponseDesc string       `json:"responseDesc"`
	Payload      []NewsDetail `json:"payload"`
}

//NewsDetail ..
type NewsDetail struct {
	ID      int
	Author  string
	Body    string
	Created string
}
