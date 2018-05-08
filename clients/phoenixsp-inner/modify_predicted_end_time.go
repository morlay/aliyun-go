package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyPredictedEndTimeRequest struct {
	requests.RpcRequest
	Caller       string `position:"Query" name:"Caller"`
	InstanceName string `position:"Query" name:"InstanceName"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	AliUID       int64  `position:"Query" name:"AliUID"`
	Bid          string `position:"Query" name:"Bid"`
	ResourceType string `position:"Query" name:"ResourceType"`
	Operator     string `position:"Query" name:"Operator"`
}

func (req *ModifyPredictedEndTimeRequest) Invoke(client *sdk.Client) (resp *ModifyPredictedEndTimeResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "ModifyPredictedEndTime", "", "")
	resp = &ModifyPredictedEndTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPredictedEndTimeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
