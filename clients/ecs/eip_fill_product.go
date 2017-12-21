package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EipFillProductRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r EipFillProductRequest) Invoke(client *sdk.Client) (response *EipFillProductResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EipFillProductRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "EipFillProduct", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		EipFillProductResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EipFillProductResponse

	err = client.DoAction(&req, &resp)
	return
}

type EipFillProductResponse struct {
	RequestId string
	Data      string
	Code      string
	Success   bool
	Message   string
}
