package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	RouteTables DescribeRouteTablesRouteTableList
}

type DescribeRouteTablesRouteTable struct {
	VRouterId      string
	RouteTableId   string
	RouteTableType string
	CreationTime   string
	RouteEntrys    DescribeRouteTablesRouteEntryList
}

type DescribeRouteTablesRouteEntry struct {
	RouteTableId         string
	DestinationCidrBlock string
	Type                 string
	Status               string
	InstanceId           string
	NextHopType          string
	NextHopRegionId      string
	NextHops             DescribeRouteTablesNextHopList
}

type DescribeRouteTablesNextHop struct {
	NextHopType     string
	NextHopId       string
	Enabled         int
	Weight          int
	NextHopRegionId string
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
