package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindInputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string `position:"Query" name:"RoleArn"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnbindInputBucketRequest) Invoke(client *sdk.Client) (response *UnbindInputBucketResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindInputBucketRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UnbindInputBucket", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindInputBucketResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindInputBucketResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindInputBucketResponse struct {
	RequestId string
}
