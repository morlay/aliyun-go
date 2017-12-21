package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (r GetStatusRequest) Invoke(client *sdk.Client) (response *GetStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "GetStatus", "cloudauth", "")

	resp := struct {
		*responses.BaseResponse
		GetStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetStatusResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetStatusData
}

type GetStatusData struct {
	StatusCode       int
	TrustedScore     float32
	SimilarityScore  float32
	AuditConclusions string
}
