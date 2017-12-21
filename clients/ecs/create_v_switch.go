package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVSwitchRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CidrBlock            string `position:"Query" name:"CidrBlock"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateVSwitchRequest) Invoke(client *sdk.Client) (response *CreateVSwitchResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateVSwitchRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateVSwitch", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateVSwitchResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateVSwitchResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateVSwitchResponse struct {
	RequestId string
	VSwitchId string
}
