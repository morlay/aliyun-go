package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	BgpGroups  DescribeBgpGroupsBgpGroupList
}

type DescribeBgpGroupsBgpGroup struct {
	Name        common.String
	Description common.String
	BgpGroupId  common.String
	PeerAsn     common.String
	AuthKey     common.String
	RouterId    common.String
	Status      common.String
	Keepalive   common.String
	LocalAsn    common.String
	Hold        common.String
	IsFake      common.String
	RouteLimit  common.String
	RegionId    common.String
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
