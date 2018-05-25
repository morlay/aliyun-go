package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (req *GetStatusRequest) Invoke(client *sdk.Client) (resp *GetStatusResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2018-05-04", "GetStatus", "cloudauth", "")
	resp = &GetStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetStatusData
}

type GetStatusData struct {
	StatusCode       common.Integer
	TrustedScore     common.Float
	SimilarityScore  common.Float
	AuditConclusions common.String
}
