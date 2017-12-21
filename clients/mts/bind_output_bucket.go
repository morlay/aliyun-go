package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindOutputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string `position:"Query" name:"RoleArn"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r BindOutputBucketRequest) Invoke(client *sdk.Client) (response *BindOutputBucketResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindOutputBucketRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "BindOutputBucket", "", "")

	resp := struct {
		*responses.BaseResponse
		BindOutputBucketResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindOutputBucketResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindOutputBucketResponse struct {
	RequestId string
}
