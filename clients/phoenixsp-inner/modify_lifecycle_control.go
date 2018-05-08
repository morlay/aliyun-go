package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyLifecycleControlRequest struct {
	requests.RpcRequest
	InstanceIDs      string `position:"Query" name:"InstanceIDs"`
	AliUID           int64  `position:"Query" name:"AliUID"`
	Bid              string `position:"Query" name:"Bid"`
	ResourceType     string `position:"Query" name:"ResourceType"`
	Operator         string `position:"Query" name:"Operator"`
	ControlLifecycle string `position:"Query" name:"ControlLifecycle"`
}

func (req *ModifyLifecycleControlRequest) Invoke(client *sdk.Client) (resp *ModifyLifecycleControlResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "ModifyLifecycleControl", "", "")
	resp = &ModifyLifecycleControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyLifecycleControlResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
