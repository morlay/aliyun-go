package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyBgpGroupAttributeRequest struct {
	requests.RpcRequest
	AuthKey              string `position:"Query" name:"AuthKey"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerAsn              int64  `position:"Query" name:"PeerAsn"`
	IsFakeAsn            string `position:"Query" name:"IsFakeAsn"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *ModifyBgpGroupAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyBgpGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBgpGroupAttribute", "vpc", "")
	resp = &ModifyBgpGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBgpGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
