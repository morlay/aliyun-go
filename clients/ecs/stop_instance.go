package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConfirmStop          string `position:"Query" name:"ConfirmStop"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	StoppedMode          string `position:"Query" name:"StoppedMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ForceStop            string `position:"Query" name:"ForceStop"`
}

func (r StopInstanceRequest) Invoke(client *sdk.Client) (response *StopInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "StopInstance", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		StopInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopInstanceResponse struct {
	RequestId string
}
