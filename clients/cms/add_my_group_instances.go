package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMyGroupInstancesRequest struct {
	requests.RpcRequest
	Instances string `position:"Query" name:"Instances"`
	GroupId   int64  `position:"Query" name:"GroupId"`
}

func (req *AddMyGroupInstancesRequest) Invoke(client *sdk.Client) (resp *AddMyGroupInstancesResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "AddMyGroupInstances", "cms", "")
	resp = &AddMyGroupInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMyGroupInstancesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
}
