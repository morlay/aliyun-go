package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMyGroupInstancesRequest struct {
	requests.RpcRequest
	Instances string `position:"Query" name:"Instances"`
	GroupId   int64  `position:"Query" name:"GroupId"`
}

func (req *UpdateMyGroupInstancesRequest) Invoke(client *sdk.Client) (resp *UpdateMyGroupInstancesResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "UpdateMyGroupInstances", "cms", "")
	resp = &UpdateMyGroupInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMyGroupInstancesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
}
