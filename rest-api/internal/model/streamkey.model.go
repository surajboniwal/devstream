package model

import "time"

type StreamKey struct {
	Id         int64      `json:"id"`
	User_Id    int64      `json:"-"`
	Name       string     `json:"name"`
	Key        int64      `json:"key"`
	Updated_At *time.Time `json:"-"`
	Created_At *time.Time `json:"created_at"`
}
