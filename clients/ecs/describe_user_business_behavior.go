package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserBusinessBehaviorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	StatusKey            string `position:"Query" name:"StatusKey"`
}

func (r DescribeUserBusinessBehaviorRequest) Invoke(client *sdk.Client) (response *DescribeUserBusinessBehaviorResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserBusinessBehaviorRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserBusinessBehavior", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserBusinessBehaviorResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserBusinessBehaviorResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserBusinessBehaviorResponse struct {
	RequestId   string
	StatusValue string
}
