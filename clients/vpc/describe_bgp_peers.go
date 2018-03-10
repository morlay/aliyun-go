package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	BgpPeers   DescribeBgpPeersBgpPeerList
}

type DescribeBgpPeersBgpPeer struct {
	Name          string
	Description   string
	BgpPeerId     string
	BgpGroupId    string
	PeerIpAddress string
	PeerAsn       int
	AuthKey       string
	RouterId      string
	BgpStatus     string
	Status        string
	Keepalive     int
	LocalAsn      int
	Hold          int
	IsFake        bool
	RouteLimit    int
	RegionId      string
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
