package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeInstallRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      string `position:"Query" name:"Force"`
	UserId     string `position:"Query" name:"UserId"`
}

func (req *NodeInstallRequest) Invoke(client *sdk.Client) (resp *NodeInstallResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeInstall", "cms", "")
	resp = &NodeInstallResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeInstallResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
}
