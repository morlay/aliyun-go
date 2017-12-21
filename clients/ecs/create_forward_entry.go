package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	InternalPort         string `position:"Query" name:"InternalPort"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ExternalIp           string `position:"Query" name:"ExternalIp"`
	ExternalPort         string `position:"Query" name:"ExternalPort"`
	InternalIp           string `position:"Query" name:"InternalIp"`
}

func (r CreateForwardEntryRequest) Invoke(client *sdk.Client) (response *CreateForwardEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateForwardEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateForwardEntry", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateForwardEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateForwardEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateForwardEntryResponse struct {
	RequestId      string
	ForwardEntryId string
}
