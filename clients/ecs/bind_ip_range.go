package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r BindIpRangeRequest) Invoke(client *sdk.Client) (response *BindIpRangeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindIpRangeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "BindIpRange", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		BindIpRangeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindIpRangeResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindIpRangeResponse struct {
	RequestId string
}
