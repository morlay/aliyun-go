package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeUserDataRequest) Invoke(client *sdk.Client) (response *DescribeUserDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserData", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserDataResponse struct {
	RequestId  string
	RegionId   string
	InstanceId string
	UserData   string
}
