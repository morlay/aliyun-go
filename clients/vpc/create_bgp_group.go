package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBgpGroupRequest struct {
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

func (r CreateBgpGroupRequest) Invoke(client *sdk.Client) (response *CreateBgpGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBgpGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBgpGroup", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateBgpGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateBgpGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateBgpGroupResponse struct {
	RequestId  string
	BgpGroupId string
}
