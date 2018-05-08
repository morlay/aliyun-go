package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetListenerAccessControlStatusRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AccessControlStatus  string `position:"Query" name:"AccessControlStatus"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetListenerAccessControlStatusRequest) Invoke(client *sdk.Client) (resp *SetListenerAccessControlStatusResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetListenerAccessControlStatus", "slb", "")
	resp = &SetListenerAccessControlStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetListenerAccessControlStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
}
