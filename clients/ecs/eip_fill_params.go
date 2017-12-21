package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EipFillParamsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r EipFillParamsRequest) Invoke(client *sdk.Client) (response *EipFillParamsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EipFillParamsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "EipFillParams", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		EipFillParamsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EipFillParamsResponse

	err = client.DoAction(&req, &resp)
	return
}

type EipFillParamsResponse struct {
	RequestId string
	Data      string
	Code      string
	Success   bool
	Message   string
}
