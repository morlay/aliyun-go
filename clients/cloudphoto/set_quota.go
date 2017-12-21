package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetQuotaRequest struct {
	requests.RpcRequest
	TotalQuota int64  `position:"Query" name:"TotalQuota"`
	LibraryId  string `position:"Query" name:"LibraryId"`
	StoreName  string `position:"Query" name:"StoreName"`
}

func (req *SetQuotaRequest) Invoke(client *sdk.Client) (resp *SetQuotaResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetQuota", "cloudphoto", "")
	resp = &SetQuotaResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetQuotaResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
