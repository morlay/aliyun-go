package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId          string
	TotalCount         int
	PageSize           int
	PageNumber         int
	ApiGroupAttributes DescribeApiGroupsApiGroupAttributeList
}

type DescribeApiGroupsApiGroupAttribute struct {
	GroupId       string
	GroupName     string
	SubDomain     string
	Description   string
	CreatedTime   string
	ModifiedTime  string
	RegionId      string
	TrafficLimit  int
	BillingStatus DescribeApiGroupsBillingStatus
	IllegalStatus DescribeApiGroupsIllegalStatus
}

type DescribeApiGroupsBillingStatus struct {
	StringValue string
}

type DescribeApiGroupsIllegalStatus struct {
	StringValue string
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
