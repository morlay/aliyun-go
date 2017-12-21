package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnbindIpRangeRequest) Invoke(client *sdk.Client) (response *UnbindIpRangeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindIpRangeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "UnbindIpRange", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		UnbindIpRangeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindIpRangeResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindIpRangeResponse struct {
	RequestId string
}
