package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InstanceLifecycleApiRequest struct {
	requests.RpcRequest
	EventResult  string `position:"Query" name:"EventResult"`
	Async        string `position:"Query" name:"Async"`
	InstanceIds  string `position:"Query" name:"InstanceIds"`
	OrderId      string `position:"Query" name:"OrderId"`
	Success      string `position:"Query" name:"Success"`
	ExtraData    string `position:"Query" name:"ExtraData"`
	EventType    string `position:"Query" name:"EventType"`
	EventTime    int64  `position:"Query" name:"EventTime"`
	AliUid       int64  `position:"Query" name:"AliUid"`
	ResourceType string `position:"Query" name:"ResourceType"`
	EventSource  string `position:"Query" name:"EventSource"`
	Token        string `position:"Query" name:"Token"`
}

func (req *InstanceLifecycleApiRequest) Invoke(client *sdk.Client) (resp *InstanceLifecycleApiResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "InstanceLifecycleApi", "", "")
	resp = &InstanceLifecycleApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type InstanceLifecycleApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
}
