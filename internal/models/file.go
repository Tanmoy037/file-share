package models

import "time"
type FileMetadata struct {
    ID        string    `json:"id"`
    Filename  string    `json:"filename"`
    Size      int64     `json:"size"`
    UploadAt  time.Time `json:"upload_at"`
    OwnerID   string    `json:"owner_id"`
}


type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
}
