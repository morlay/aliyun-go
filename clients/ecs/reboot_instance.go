package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RebootInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ForceStop            string `position:"Query" name:"ForceStop"`
}

func (r RebootInstanceRequest) Invoke(client *sdk.Client) (response *RebootInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RebootInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "RebootInstance", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		RebootInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RebootInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RebootInstanceResponse struct {
	RequestId string
}
