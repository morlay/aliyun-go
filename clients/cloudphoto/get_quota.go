package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetQuotaRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetQuotaRequest) Invoke(client *sdk.Client) (resp *GetQuotaResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetQuota", "cloudphoto", "")
	resp = &GetQuotaResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQuotaResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Quota     GetQuotaQuota
}

type GetQuotaQuota struct {
	TotalQuota  common.Long
	FacesCount  common.Integer
	PhotosCount common.Integer
	UsedQuota   common.Long
	VideosCount common.Integer
}
