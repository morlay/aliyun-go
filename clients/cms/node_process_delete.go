package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeProcessDeleteRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	Id         string `position:"Query" name:"Id"`
}

func (r NodeProcessDeleteRequest) Invoke(client *sdk.Client) (response *NodeProcessDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeProcessDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeProcessDelete", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeProcessDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeProcessDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeProcessDeleteResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
