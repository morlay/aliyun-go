package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetQuotaRequest struct {
	TotalQuota int64  `position:"Query" name:"TotalQuota"`
	LibraryId  string `position:"Query" name:"LibraryId"`
	StoreName  string `position:"Query" name:"StoreName"`
}

func (r SetQuotaRequest) Invoke(client *sdk.Client) (response *SetQuotaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetQuotaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetQuota", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		SetQuotaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetQuotaResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetQuotaResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
