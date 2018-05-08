package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiGroupsRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	GroupName  string `position:"Query" name:"GroupName"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeApiGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeApiGroupsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiGroups", "apigateway", "")
	resp = &DescribeApiGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiGroupsResponse struct {
	responses.BaseResponse
	RequestId          common.String
	TotalCount         common.Integer
	PageSize           common.Integer
	PageNumber         common.Integer
	ApiGroupAttributes DescribeApiGroupsApiGroupAttributeList
}

type DescribeApiGroupsApiGroupAttribute struct {
	GroupId       common.String
	GroupName     common.String
	SubDomain     common.String
	Description   common.String
	CreatedTime   common.String
	ModifiedTime  common.String
	RegionId      common.String
	TrafficLimit  common.Integer
	BillingStatus DescribeApiGroupsBillingStatus
	IllegalStatus DescribeApiGroupsIllegalStatus
}

type DescribeApiGroupsBillingStatus struct {
	StringValue common.String
}

type DescribeApiGroupsIllegalStatus struct {
	StringValue common.String
}

type DescribeApiGroupsApiGroupAttributeList []DescribeApiGroupsApiGroupAttribute

func (list *DescribeApiGroupsApiGroupAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupsApiGroupAttribute)
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
