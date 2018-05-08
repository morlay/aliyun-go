package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BindOutputBucketRequest struct {
	requests.RpcRequest
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string `position:"Query" name:"RoleArn"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *BindOutputBucketRequest) Invoke(client *sdk.Client) (resp *BindOutputBucketResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "BindOutputBucket", "mts", "")
	resp = &BindOutputBucketResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindOutputBucketResponse struct {
	responses.BaseResponse
	RequestId common.String
}
