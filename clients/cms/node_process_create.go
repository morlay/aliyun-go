package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeProcessCreate", "cms", "")
	resp = &NodeProcessCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeProcessCreateResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
