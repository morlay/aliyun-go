package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	RequestId string
	Action    string
	Library   GetLibraryLibrary
}

type GetLibraryLibrary struct {
	Quota           GetLibraryQuota
	AutoCleanConfig GetLibraryAutoCleanConfig
}

type GetLibraryQuota struct {
	TotalQuota      int64
	TotalTrashQuota int64
	FacesCount      int
	PhotosCount     int
	UsedQuota       int64
	VideosCount     int
	ActiveSize      int64
	InactiveSize    int64
}

type GetLibraryAutoCleanConfig struct {
	AutoCleanEnabled bool
	AutoCleanDays    int
}
