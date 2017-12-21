package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeUninstallRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r NodeUninstallRequest) Invoke(client *sdk.Client) (response *NodeUninstallResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeUninstallRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeUninstall", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeUninstallResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeUninstallResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeUninstallResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
