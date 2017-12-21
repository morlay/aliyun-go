package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AddIpRangeRequest) Invoke(client *sdk.Client) (response *AddIpRangeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddIpRangeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AddIpRange", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AddIpRangeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddIpRangeResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddIpRangeResponse struct {
	RequestId string
}
