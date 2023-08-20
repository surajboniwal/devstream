package model

type StreamKey struct {
	User_Id int64  `json:"-"`
	Name    string `json:"name"`
	Key     int64  `json:"key"`
}
