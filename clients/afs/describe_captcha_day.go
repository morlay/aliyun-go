package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCaptchaDayRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
}

func (req *DescribeCaptchaDayRequest) Invoke(client *sdk.Client) (resp *DescribeCaptchaDayResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaDay", "", "")
	resp = &DescribeCaptchaDayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCaptchaDayResponse struct {
	responses.BaseResponse
	RequestId  string
	BizCode    string
	HasData    bool
	CaptchaDay DescribeCaptchaDayCaptchaDay
}

type DescribeCaptchaDayCaptchaDay struct {
	Init                        int
	AskForVerify                int
	DirecetStrategyInterception int
	TwiceVerify                 int
	Pass                        int
	CheckTested                 int
	UncheckTested               int
	LegalSign                   int
	MaliciousFlow               int
}
