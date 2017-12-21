package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteForwardEntryRequest) Invoke(client *sdk.Client) (response *DeleteForwardEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteForwardEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteForwardEntry", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteForwardEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteForwardEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteForwardEntryResponse struct {
	RequestId string
}
