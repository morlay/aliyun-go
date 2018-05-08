package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBgpPeerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpPeerId            string `position:"Query" name:"BgpPeerId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteBgpPeerRequest) Invoke(client *sdk.Client) (resp *DeleteBgpPeerResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpPeer", "vpc", "")
	resp = &DeleteBgpPeerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBgpPeerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
