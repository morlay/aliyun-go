package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRouteTablesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Type                 string `position:"Query" name:"Type"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	RouterType           string `position:"Query" name:"RouterType"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	RouterId             string `position:"Query" name:"RouterId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
}

func (req *DescribeRouteTablesRequest) Invoke(client *sdk.Client) (resp *DescribeRouteTablesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteTables", "vpc", "")
	resp = &DescribeRouteTablesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRouteTablesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	PageNumber  common.Integer
	PageSize    common.Integer
	RouteTables DescribeRouteTablesRouteTableList
}

type DescribeRouteTablesRouteTable struct {
	VRouterId      common.String
	RouteTableId   common.String
	RouteTableType common.String
	CreationTime   common.String
	RouteEntrys    DescribeRouteTablesRouteEntryList
}

type DescribeRouteTablesRouteEntry struct {
	RouteTableId         common.String
	DestinationCidrBlock common.String
	Type                 common.String
	Status               common.String
	InstanceId           common.String
	NextHopType          common.String
	NextHopRegionId      common.String
	NextHops             DescribeRouteTablesNextHopList
}

type DescribeRouteTablesNextHop struct {
	NextHopType     common.String
	NextHopId       common.String
	Enabled         common.Integer
	Weight          common.Integer
	NextHopRegionId common.String
}

type DescribeRouteTablesRouteTableList []DescribeRouteTablesRouteTable

func (list *DescribeRouteTablesRouteTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteTable)
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

type DescribeRouteTablesRouteEntryList []DescribeRouteTablesRouteEntry

func (list *DescribeRouteTablesRouteEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteEntry)
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

type DescribeRouteTablesNextHopList []DescribeRouteTablesNextHop

func (list *DescribeRouteTablesNextHopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesNextHop)
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
