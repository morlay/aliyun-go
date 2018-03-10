package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhotoRequest struct {
	requests.RpcRequest
	PhotoTitle      string `position:"Query" name:"PhotoTitle"`
	LibraryId       string `position:"Query" name:"LibraryId"`
	ShareExpireTime int64  `position:"Query" name:"ShareExpireTime"`
	StoreName       string `position:"Query" name:"StoreName"`
	UploadType      string `position:"Query" name:"UploadType"`
	Remark          string `position:"Query" name:"Remark"`
	SessionId       string `position:"Query" name:"SessionId"`
	Staging         string `position:"Query" name:"Staging"`
	FileId          string `position:"Query" name:"FileId"`
}

func (req *CreatePhotoRequest) Invoke(client *sdk.Client) (resp *CreatePhotoResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhoto", "cloudphoto", "")
	resp = &CreatePhotoResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePhotoResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Photo     CreatePhotoPhoto
}

type CreatePhotoPhoto struct {
	Id              int64
	IdStr           string
	Title           string
	FileId          string
	Location        string
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
