package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhotoRequest struct {
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

func (r CreatePhotoRequest) Invoke(client *sdk.Client) (response *CreatePhotoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePhotoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhoto", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		CreatePhotoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePhotoResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Photo     CreatePhotoPhoto
}

type CreatePhotoPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Remark          string
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
}
