package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceVncPasswdRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VncPassword          string `position:"Query" name:"VncPassword"`
}

func (r ModifyInstanceVncPasswdRequest) Invoke(client *sdk.Client) (response *ModifyInstanceVncPasswdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceVncPasswdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceVncPasswd", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceVncPasswdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceVncPasswdResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceVncPasswdResponse struct {
	RequestId string
}
