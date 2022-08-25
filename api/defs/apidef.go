package defs

// requests
type UserCredential struct {
	Username string `json: "user_name"`
	Pwd string `json: "pwd"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

// response
type SignedUp struct {
	Success bool `json: "success"`
	SessionId string `json: "session_id"`
}

type UserSession struct {
	Username string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

// data model
type VideoInfo struct {
	Id string
	AuthorId int 
	Name string 
	DisplayCtime string
}

type Comment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

type SimpleSession struct {
	Username string // login name
	TTL int64
}

type User struct {
	Id int
	LoginName string
	Pwd string
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}