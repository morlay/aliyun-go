package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegisterPhotoRequest struct {
	requests.RpcRequest
	LibraryId  string  `position:"Query" name:"LibraryId"`
	Latitude   float32 `position:"Query" name:"Latitude"`
	PhotoTitle string  `position:"Query" name:"PhotoTitle"`
	StoreName  string  `position:"Query" name:"StoreName"`
	IsVideo    string  `position:"Query" name:"IsVideo"`
	Remark     string  `position:"Query" name:"Remark"`
	Size       int64   `position:"Query" name:"Size"`
	TakenAt    int64   `position:"Query" name:"TakenAt"`
	Width      int     `position:"Query" name:"Width"`
	Location   string  `position:"Query" name:"Location"`
	Longitude  float32 `position:"Query" name:"Longitude"`
	Height     int     `position:"Query" name:"Height"`
	Md5        string  `position:"Query" name:"Md.5"`
}

func (req *RegisterPhotoRequest) Invoke(client *sdk.Client) (resp *RegisterPhotoResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RegisterPhoto", "cloudphoto", "")
	resp = &RegisterPhotoResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterPhotoResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Photo     RegisterPhotoPhoto
}

type RegisterPhotoPhoto struct {
	Id              int64
	IdStr           string
	Title           string
	Location        string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Size            int64
	Remark          string
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
}
