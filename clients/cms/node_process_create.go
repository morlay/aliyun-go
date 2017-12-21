package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeProcessCreateRequest struct {
	InstanceId  string `position:"Query" name:"InstanceId"`
	ProcessName string `position:"Query" name:"ProcessName"`
	Name        string `position:"Query" name:"Name"`
	ProcessUser string `position:"Query" name:"ProcessUser"`
	Command     string `position:"Query" name:"Command"`
}

func (r NodeProcessCreateRequest) Invoke(client *sdk.Client) (response *NodeProcessCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeProcessCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeProcessCreate", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeProcessCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeProcessCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeProcessCreateResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
