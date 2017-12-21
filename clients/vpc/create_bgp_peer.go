package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBgpPeerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerIpAddress        string `position:"Query" name:"PeerIpAddress"`
}

func (r CreateBgpPeerRequest) Invoke(client *sdk.Client) (response *CreateBgpPeerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBgpPeerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBgpPeer", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateBgpPeerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateBgpPeerResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateBgpPeerResponse struct {
	RequestId string
	BgpPeerId string
}
