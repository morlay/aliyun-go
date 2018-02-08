package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCaptchaRiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
}

func (req *DescribeCaptchaRiskRequest) Invoke(client *sdk.Client) (resp *DescribeCaptchaRiskResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaRisk", "", "")
	resp = &DescribeCaptchaRiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCaptchaRiskResponse struct {
	responses.BaseResponse
	RequestId      string
	BizCode        string
	NumOfThisMonth int
	NumOfLastMonth int
	RiskLevel      string
}
