package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeInstallRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      string `position:"Query" name:"Force"`
	UserId     string `position:"Query" name:"UserId"`
}

func (req *NodeInstallRequest) Invoke(client *sdk.Client) (resp *NodeInstallResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeInstall", "cms", "")
	resp = &NodeInstallResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeInstallResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
