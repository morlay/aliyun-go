package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeStatusRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r NodeStatusRequest) Invoke(client *sdk.Client) (response *NodeStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeStatus", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeStatusResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	InstanceId   string
	AutoInstall  bool
	Status       string
}
