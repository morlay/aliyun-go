package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreatePhotoRequest struct {
	requests.RpcRequest
	TakenAt         int64  `position:"Query" name:"TakenAt"`
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Photo     CreatePhotoPhoto
}

type CreatePhotoPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	FileId          common.String
	Location        common.String
	State           common.String
	Md5             common.String
	IsVideo         bool
	Size            common.Long
	Remark          common.String
	Width           common.Long
	Height          common.Long
	Ctime           common.Long
	Mtime           common.Long
	TakenAt         common.Long
	ShareExpireTime common.Long
}
