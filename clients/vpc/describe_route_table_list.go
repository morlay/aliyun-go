package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRouteTableListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	RouterType           string `position:"Query" name:"RouterType"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
}

func (req *DescribeRouteTableListRequest) Invoke(client *sdk.Client) (resp *DescribeRouteTableListResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteTableList", "vpc", "")
	resp = &DescribeRouteTableListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRouteTableListResponse struct {
	responses.BaseResponse
	RequestId       common.String
	Code            common.String
	Message         common.String
	Success         bool
	PageSize        common.Integer
	PageNumber      common.Integer
	TotalCount      common.Integer
	RouterTableList DescribeRouteTableListRouterTableListTypeList
}

type DescribeRouteTableListRouterTableListType struct {
	VpcId          common.String
	RouterType     common.String
	RouterId       common.String
	RouteTableId   common.String
	RouteTableName common.String
	RouteTableType common.String
	Description    common.String
	CreationTime   common.String
}

type DescribeRouteTableListRouterTableListTypeList []DescribeRouteTableListRouterTableListType

func (list *DescribeRouteTableListRouterTableListTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTableListRouterTableListType)
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
