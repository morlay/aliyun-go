package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddListenerWhiteListItemRequest struct {
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

func (req *AddListenerWhiteListItemRequest) Invoke(client *sdk.Client) (resp *AddListenerWhiteListItemResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "AddListenerWhiteListItem", "slb", "")
	resp = &AddListenerWhiteListItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddListenerWhiteListItemResponse struct {
	responses.BaseResponse
	RequestId common.String
}
