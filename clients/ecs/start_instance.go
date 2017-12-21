package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartInstanceRequest struct {
	InitLocalDisk        string `position:"Query" name:"InitLocalDisk"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r StartInstanceRequest) Invoke(client *sdk.Client) (response *StartInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "StartInstance", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		StartInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartInstanceResponse struct {
	RequestId string
}
