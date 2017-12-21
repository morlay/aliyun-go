package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveListenerWhiteListItemRequest struct {
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

func (r RemoveListenerWhiteListItemRequest) Invoke(client *sdk.Client) (response *RemoveListenerWhiteListItemResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveListenerWhiteListItemRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveListenerWhiteListItem", "slb", "")

	resp := struct {
		*responses.BaseResponse
		RemoveListenerWhiteListItemResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveListenerWhiteListItemResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveListenerWhiteListItemResponse struct {
	RequestId string
}
