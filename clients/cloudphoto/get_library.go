package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetLibraryRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetLibraryRequest) Invoke(client *sdk.Client) (resp *GetLibraryResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetLibrary", "cloudphoto", "")
	resp = &GetLibraryResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLibraryResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Library   GetLibraryLibrary
}

type GetLibraryLibrary struct {
	Ctime           common.Long
	Quota           GetLibraryQuota
	AutoCleanConfig GetLibraryAutoCleanConfig
}

type GetLibraryQuota struct {
	TotalQuota      common.Long
	TotalTrashQuota common.Long
	FacesCount      common.Integer
	PhotosCount     common.Integer
	UsedQuota       common.Long
	VideosCount     common.Integer
	ActiveSize      common.Long
	InactiveSize    common.Long
}

type GetLibraryAutoCleanConfig struct {
	AutoCleanEnabled bool
	AutoCleanDays    common.Integer
}
