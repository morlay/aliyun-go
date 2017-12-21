package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBgpPeersRequest struct {
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

func (r DescribeBgpPeersRequest) Invoke(client *sdk.Client) (response *DescribeBgpPeersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBgpPeersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBgpPeers", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeBgpPeersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeBgpPeersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeBgpPeersResponse struct {
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
	PeerAsn       string
	AuthKey       string
	RouterId      string
	BgpStatus     string
	Status        string
	Keepalive     string
	LocalAsn      string
	Hold          string
	IsFake        string
	RouteLimit    string
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
