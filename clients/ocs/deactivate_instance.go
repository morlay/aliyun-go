package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateInstanceRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (req *DeactivateInstanceRequest) Invoke(client *sdk.Client) (resp *DeactivateInstanceResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DeactivateInstance", "", "")
	resp = &DeactivateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeactivateInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
