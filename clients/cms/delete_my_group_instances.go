package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteMyGroupInstancesRequest struct {
	requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
	GroupId     int64  `position:"Query" name:"GroupId"`
}

func (req *DeleteMyGroupInstancesRequest) Invoke(client *sdk.Client) (resp *DeleteMyGroupInstancesResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteMyGroupInstances", "cms", "")
	resp = &DeleteMyGroupInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMyGroupInstancesResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
}
