package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeIpControlPolicyItemsRequest struct {
	requests.RpcRequest
	IpControlId  string `position:"Query" name:"IpControlId"`
	PolicyItemId string `position:"Query" name:"PolicyItemId"`
	PageSize     int    `position:"Query" name:"PageSize"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeIpControlPolicyItemsRequest) Invoke(client *sdk.Client) (resp *DescribeIpControlPolicyItemsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeIpControlPolicyItems", "apigateway", "")
	resp = &DescribeIpControlPolicyItemsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeIpControlPolicyItemsResponse struct {
	responses.BaseResponse
	RequestId            string
	TotalCount           int
	PageSize             int
	PageNumber           int
	IpControlPolicyItems DescribeIpControlPolicyItemsIpControlPolicyItemList
}

type DescribeIpControlPolicyItemsIpControlPolicyItem struct {
	AppId        string
	CidrIp       string
	PolicyItemId string
	CreateTime   string
	ModifiedTime string
}

type DescribeIpControlPolicyItemsIpControlPolicyItemList []DescribeIpControlPolicyItemsIpControlPolicyItem

func (list *DescribeIpControlPolicyItemsIpControlPolicyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeIpControlPolicyItemsIpControlPolicyItem)
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
