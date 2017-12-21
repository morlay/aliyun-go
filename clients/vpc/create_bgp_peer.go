package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBgpPeerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerIpAddress        string `position:"Query" name:"PeerIpAddress"`
}

func (req *CreateBgpPeerRequest) Invoke(client *sdk.Client) (resp *CreateBgpPeerResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBgpPeer", "vpc", "")
	resp = &CreateBgpPeerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBgpPeerResponse struct {
	responses.BaseResponse
	RequestId string
	BgpPeerId string
}
