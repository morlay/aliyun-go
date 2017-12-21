package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeInstallRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      string `position:"Query" name:"Force"`
	UserId     string `position:"Query" name:"UserId"`
}

func (r NodeInstallRequest) Invoke(client *sdk.Client) (response *NodeInstallResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeInstallRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeInstall", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeInstallResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeInstallResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeInstallResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
