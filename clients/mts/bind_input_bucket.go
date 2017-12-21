package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindInputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string `position:"Query" name:"RoleArn"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r BindInputBucketRequest) Invoke(client *sdk.Client) (response *BindInputBucketResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindInputBucketRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "BindInputBucket", "", "")

	resp := struct {
		*responses.BaseResponse
		BindInputBucketResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindInputBucketResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindInputBucketResponse struct {
	RequestId string
}
