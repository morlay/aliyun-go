package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLibraryRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r GetLibraryRequest) Invoke(client *sdk.Client) (response *GetLibraryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetLibraryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetLibrary", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetLibraryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetLibraryResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetLibraryResponse struct {
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
	TotalQuota  int64
	FacesCount  int
	PhotosCount int
	UsedQuota   int64
	VideosCount int
}

type GetLibraryAutoCleanConfig struct {
	AutoCleanEnabled bool
	AutoCleanDays    int
}
