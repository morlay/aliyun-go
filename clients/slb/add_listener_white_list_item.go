package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddListenerWhiteListItemRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	SourceItems          string `position:"Query" name:"SourceItems"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r AddListenerWhiteListItemRequest) Invoke(client *sdk.Client) (response *AddListenerWhiteListItemResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddListenerWhiteListItemRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "AddListenerWhiteListItem", "slb", "")

	resp := struct {
		*responses.BaseResponse
		AddListenerWhiteListItemResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddListenerWhiteListItemResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddListenerWhiteListItemResponse struct {
	RequestId string
}
