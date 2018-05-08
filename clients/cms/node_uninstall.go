package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeUninstallRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *NodeUninstallRequest) Invoke(client *sdk.Client) (resp *NodeUninstallResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeUninstall", "cms", "")
	resp = &NodeUninstallResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeUninstallResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
}
