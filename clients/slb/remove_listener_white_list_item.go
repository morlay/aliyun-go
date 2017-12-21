package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveListenerWhiteListItemRequest struct {
	requests.RpcRequest
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

func (req *RemoveListenerWhiteListItemRequest) Invoke(client *sdk.Client) (resp *RemoveListenerWhiteListItemResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveListenerWhiteListItem", "slb", "")
	resp = &RemoveListenerWhiteListItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveListenerWhiteListItemResponse struct {
	responses.BaseResponse
	RequestId string
}
