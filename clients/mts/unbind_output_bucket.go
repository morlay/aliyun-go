package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindOutputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnbindOutputBucketRequest) Invoke(client *sdk.Client) (response *UnbindOutputBucketResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindOutputBucketRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UnbindOutputBucket", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindOutputBucketResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindOutputBucketResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindOutputBucketResponse struct {
	RequestId string
}
