package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	BizCode        common.String
	NumOfThisMonth common.Integer
	NumOfLastMonth common.Integer
	RiskLevel      common.String
}
