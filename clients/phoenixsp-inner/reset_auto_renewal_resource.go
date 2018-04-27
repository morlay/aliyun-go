package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAutoRenewalResourceRequest struct {
	requests.RpcRequest
	InstanceNames         string `position:"Query" name:"InstanceNames"`
	Caller                string `position:"Query" name:"Caller"`
	RenewalDuration       int    `position:"Query" name:"RenewalDuration"`
	AutoRenewal           string `position:"Query" name:"AutoRenewal"`
	RenewalStatus         string `position:"Query" name:"RenewalStatus"`
	SaleCycle             string `position:"Query" name:"SaleCycle"`
	AutoRenewalOffsetDays string `position:"Query" name:"AutoRenewalOffsetDays"`
	AliUID                int64  `position:"Query" name:"AliUID"`
	RenewalCycUnit        int    `position:"Query" name:"RenewalCycUnit"`
	Bid                   string `position:"Query" name:"Bid"`
	ResourceType          string `position:"Query" name:"ResourceType"`
	Operator              string `position:"Query" name:"Operator"`
}

func (req *ResetAutoRenewalResourceRequest) Invoke(client *sdk.Client) (resp *ResetAutoRenewalResourceResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "ResetAutoRenewalResource", "", "")
	resp = &ResetAutoRenewalResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAutoRenewalResourceResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
}
