package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetQuotaRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r GetQuotaRequest) Invoke(client *sdk.Client) (response *GetQuotaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetQuotaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetQuota", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetQuotaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetQuotaResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetQuotaResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Quota     GetQuotaQuota
}

type GetQuotaQuota struct {
	TotalQuota  int64
	FacesCount  int
	PhotosCount int
	UsedQuota   int64
	VideosCount int
}
