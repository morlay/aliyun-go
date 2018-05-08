package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBgpPeersRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	BgpPeerId            string `position:"Query" name:"BgpPeerId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeBgpPeersRequest) Invoke(client *sdk.Client) (resp *DescribeBgpPeersResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBgpPeers", "vpc", "")
	resp = &DescribeBgpPeersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBgpPeersResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	BgpPeers   DescribeBgpPeersBgpPeerList
}

type DescribeBgpPeersBgpPeer struct {
	Name          common.String
	Description   common.String
	BgpPeerId     common.String
	BgpGroupId    common.String
	PeerIpAddress common.String
	PeerAsn       common.String
	AuthKey       common.String
	RouterId      common.String
	BgpStatus     common.String
	Status        common.String
	Keepalive     common.String
	LocalAsn      common.String
	Hold          common.String
	IsFake        common.String
	RouteLimit    common.String
	RegionId      common.String
}

type DescribeBgpPeersBgpPeerList []DescribeBgpPeersBgpPeer

func (list *DescribeBgpPeersBgpPeerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpPeersBgpPeer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
