package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeProcessDeleteRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	Id         string `position:"Query" name:"Id"`
}

func (req *NodeProcessDeleteRequest) Invoke(client *sdk.Client) (resp *NodeProcessDeleteResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeProcessDelete", "cms", "")
	resp = &NodeProcessDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeProcessDeleteResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
}
