package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternalIp           string `position:"Query" name:"InternalIp"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	InternalPort         string `position:"Query" name:"InternalPort"`
	ExternalIp           string `position:"Query" name:"ExternalIp"`
	ExternalPort         string `position:"Query" name:"ExternalPort"`
}

func (r ModifyForwardEntryRequest) Invoke(client *sdk.Client) (response *ModifyForwardEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyForwardEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyForwardEntry", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyForwardEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyForwardEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyForwardEntryResponse struct {
	RequestId string
}
