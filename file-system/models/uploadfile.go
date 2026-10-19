package models

type UploadFile struct {
	Message  string `json:"message"`
	FileName string `json:"filename"`
	Size     int64  `json:"size"`
	// Type     string `json:"type"`
}
