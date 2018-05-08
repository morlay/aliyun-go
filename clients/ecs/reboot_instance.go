package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RebootInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ForceStop            string `position:"Query" name:"ForceStop"`
}

func (req *RebootInstanceRequest) Invoke(client *sdk.Client) (resp *RebootInstanceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "RebootInstance", "ecs", "")
	resp = &RebootInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RebootInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
