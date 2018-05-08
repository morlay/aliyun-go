package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeStatusRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *NodeStatusRequest) Invoke(client *sdk.Client) (resp *NodeStatusResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeStatus", "cms", "")
	resp = &NodeStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeStatusResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
	InstanceId   common.String
	AutoInstall  bool
	Status       common.String
}
