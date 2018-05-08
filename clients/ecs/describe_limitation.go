package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLimitationRequest struct {
	requests.RpcRequest
	Limitation           string `position:"Query" name:"Limitation"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLimitationRequest) Invoke(client *sdk.Client) (resp *DescribeLimitationResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeLimitation", "ecs", "")
	resp = &DescribeLimitationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLimitationResponse struct {
	responses.BaseResponse
	RequestId  common.String
	Limitation common.String
	Value      common.String
}
