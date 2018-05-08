package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeProcessCreateRequest struct {
	requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	ProcessName string `position:"Query" name:"ProcessName"`
	Name        string `position:"Query" name:"Name"`
	ProcessUser string `position:"Query" name:"ProcessUser"`
	Command     string `position:"Query" name:"Command"`
}

func (req *NodeProcessCreateRequest) Invoke(client *sdk.Client) (resp *NodeProcessCreateResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeProcessCreate", "cms", "")
	resp = &NodeProcessCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeProcessCreateResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
}
