package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVSwitchAttributeRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyVSwitchAttributeRequest) Invoke(client *sdk.Client) (response *ModifyVSwitchAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyVSwitchAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVSwitchAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyVSwitchAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyVSwitchAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyVSwitchAttributeResponse struct {
	RequestId string
}
