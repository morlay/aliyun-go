package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLimitationRequest struct {
	Limitation           string `position:"Query" name:"Limitation"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLimitationRequest) Invoke(client *sdk.Client) (response *DescribeLimitationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLimitationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeLimitation", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLimitationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLimitationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLimitationResponse struct {
	RequestId  string
	Limitation string
	Value      string
}
