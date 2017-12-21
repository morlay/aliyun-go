package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBgpPeerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpPeerId            string `position:"Query" name:"BgpPeerId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteBgpPeerRequest) Invoke(client *sdk.Client) (response *DeleteBgpPeerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBgpPeerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpPeer", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBgpPeerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBgpPeerResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBgpPeerResponse struct {
	RequestId string
}
