package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBgpGroupsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeBgpGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeBgpGroupsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBgpGroups", "vpc", "")
	resp = &DescribeBgpGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBgpGroupsResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	BgpGroups  DescribeBgpGroupsBgpGroupList
}

type DescribeBgpGroupsBgpGroup struct {
	Name        string
	Description string
	BgpGroupId  string
	PeerAsn     int
	AuthKey     string
	RouterId    string
	Status      string
	Keepalive   int
	LocalAsn    int
	Hold        int
	IsFake      bool
	RouteLimit  int
	RegionId    string
}

type DescribeBgpGroupsBgpGroupList []DescribeBgpGroupsBgpGroup

func (list *DescribeBgpGroupsBgpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpGroupsBgpGroup)
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
