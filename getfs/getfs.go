package getfs

import (
	"encoding/json"
	"errors"

	resty "github.com/go-resty/resty/v2"
)

const (
	url   = "https://api.fshare.vn/api"
	appID = "L2S7R6ZMagggC5wWkQhX2+aDi467PPuftWUMRFSn"
)

// Service represent for general service
type Service struct {
	UserEmail string `json:"user_email"`
	Password  string `json:"password"`
	AppID     string `json:"app_id"`
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
}

// LoginResp for login response
type LoginResp struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
}

// GetLinkResp for get link response
type GetLinkResp struct {
	Location string `json:"location"`
}

// GetFileInfoResp for get file info response
type GetFileInfoResp struct {
	ID            string `json:"id"`
	LinkCode      string `json:"linkcode"`
	Name          string `json:"name"`
	Shared        string `json:"shared"`
	DirectLink    string `json:"directlink"`
	Secure        string `json:"secure"`
	HashIndex     string `json:"hash_index"`
	Public        string `json:"public"`
	Path          string `json:"path"`
	OwnerID       string `json:"owner_id"`
	Size          string `json:"size"`
	DownloadCount string `json:"downloadcount"`
	Deleted       string `json:"deleted"`
	Description   string `json:"description"`
	Created       string `json:"created"`
	Modified      string `json:"modified"`
	MimeType      string `json:"mimetype"`
	FileType      string `json:"file_type"`
	PWD           string `json:"pwd"`
	PID           string `json:"pid"`
	CRC32         string `json:"crc32"`
	FolderPath    string `json:"folder_path"`
	StorageID     string `json:"42"`
	RealName      string `json:"realname"`
}

// NewService represent for new Service
func NewService(userEmail, password string) *Service {
	return &Service{
		UserEmail: userEmail,
		Password:  password,
	}
}

// Login to fshare
func (s *Service) Login() error {
	resty := resty.New()
	// resty.SetDebug(true)
	resty.SetHostURL(url)
	resp, err := resty.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
	}).SetBody(map[string]interface{}{
		"user_email": s.UserEmail,
		"password":   s.Password,
		"app_key":    appID,
	}).Post("/user/login")
	if err != nil {
		return err
	}

	data := new(LoginResp)
	json.Unmarshal(resp.Body(), &data)

	if resp.StatusCode() == 200 {
		s.SessionID = data.SessionID
		s.Token = data.Token
	} else {
		return errors.New(data.Msg)
	}

	return nil
}

// GetFileInfo of file on fshare by url
func (s *Service) GetFileInfo(fileURL string) (*GetFileInfoResp, error) {
	resty := resty.New()
	// resty.SetDebug(true)
	resty.SetHostURL(url)
	resp, err := resty.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
	}).SetBody(map[string]interface{}{
		"token": s.Token,
		"url":   fileURL,
	}).Post("/fileops/get")
	if err != nil {
		return nil, err
	}

	data := new(GetFileInfoResp)
	json.Unmarshal(resp.Body(), &data)

	return data, nil
}

// GetLink of file on fshare by url
func (s *Service) GetLink(fileURL string) (*GetLinkResp, error) {
	resty := resty.New()
	// resty.SetDebug(true)
	resty.SetHostURL(url)
	resp, err := resty.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
	}).SetBody(map[string]interface{}{
		"token": s.Token,
		"url":   fileURL,
	}).Post("/session/download")
	if err != nil {
		return nil, err
	}

	data := new(GetLinkResp)
	json.Unmarshal(resp.Body(), &data)

	return data, nil
}
