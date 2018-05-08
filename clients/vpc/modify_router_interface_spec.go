package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyRouterInterfaceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
}

func (req *ModifyRouterInterfaceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyRouterInterfaceSpecResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouterInterfaceSpec", "vpc", "")
	resp = &ModifyRouterInterfaceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyRouterInterfaceSpecResponse struct {
	responses.BaseResponse
	RequestId common.String
	Spec      common.String
}
