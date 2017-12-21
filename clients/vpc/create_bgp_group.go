package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBgpGroupRequest struct {
	requests.RpcRequest
	AuthKey              string `position:"Query" name:"AuthKey"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerAsn              int64  `position:"Query" name:"PeerAsn"`
	IsFakeAsn            string `position:"Query" name:"IsFakeAsn"`
	RouterId             string `position:"Query" name:"RouterId"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *CreateBgpGroupRequest) Invoke(client *sdk.Client) (resp *CreateBgpGroupResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBgpGroup", "vpc", "")
	resp = &CreateBgpGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBgpGroupResponse struct {
	responses.BaseResponse
	RequestId  string
	BgpGroupId string
}
