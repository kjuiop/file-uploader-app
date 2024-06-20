package models

type FileUploadReq struct {
	CustomerId string `json:"customer_id"`
	FileType   string `json:"file_type"`
}

func (frq *FileUploadReq) CheckValid() error {
	return nil
}
