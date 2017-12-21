package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserBusinessBehaviorRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	StatusKey            string `position:"Query" name:"StatusKey"`
}

func (req *DescribeUserBusinessBehaviorRequest) Invoke(client *sdk.Client) (resp *DescribeUserBusinessBehaviorResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserBusinessBehavior", "ecs", "")
	resp = &DescribeUserBusinessBehaviorResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserBusinessBehaviorResponse struct {
	responses.BaseResponse
	RequestId   string
	StatusValue string
}
