package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RiskControlRequest struct {
	requests.RpcRequest
	Reason        string `position:"Query" name:"Reason"`
	RiskCondition string `position:"Query" name:"RiskCondition"`
	Aliuid        int64  `position:"Query" name:"Aliuid"`
	Operation     string `position:"Query" name:"Operation"`
	Token         string `position:"Query" name:"Token"`
}

func (req *RiskControlRequest) Invoke(client *sdk.Client) (resp *RiskControlResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "RiskControl", "", "")
	resp = &RiskControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type RiskControlResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
}
