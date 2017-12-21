package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeStatusRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *NodeStatusRequest) Invoke(client *sdk.Client) (resp *NodeStatusResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeStatus", "cms", "")
	resp = &NodeStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeStatusResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	InstanceId   string
	AutoInstall  bool
	Status       string
}
