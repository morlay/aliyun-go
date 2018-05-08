package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTrafficControlsRequest struct {
	requests.RpcRequest
	TrafficControlId   string `position:"Query" name:"TrafficControlId"`
	GroupId            string `position:"Query" name:"GroupId"`
	ApiId              string `position:"Query" name:"ApiId"`
	StageName          string `position:"Query" name:"StageName"`
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	PageNumber         int    `position:"Query" name:"PageNumber"`
	PageSize           int    `position:"Query" name:"PageSize"`
}

func (req *DescribeTrafficControlsRequest) Invoke(client *sdk.Client) (resp *DescribeTrafficControlsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeTrafficControls", "apigateway", "")
	resp = &DescribeTrafficControlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTrafficControlsResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TotalCount      common.Integer
	PageSize        common.Integer
	PageNumber      common.Integer
	TrafficControls DescribeTrafficControlsTrafficControlList
}

type DescribeTrafficControlsTrafficControl struct {
	TrafficControlId   common.String
	TrafficControlName common.String
	Description        common.String
	TrafficControlUnit common.String
	ApiDefault         common.Integer
	UserDefault        common.Integer
	AppDefault         common.Integer
	CreatedTime        common.String
	ModifiedTime       common.String
	SpecialPolicies    DescribeTrafficControlsSpecialPolicyList
}

type DescribeTrafficControlsSpecialPolicy struct {
	SpecialType common.String
	Specials    DescribeTrafficControlsSpecialList
}

type DescribeTrafficControlsSpecial struct {
	SpecialKey   common.String
	TrafficValue common.Integer
}

type DescribeTrafficControlsTrafficControlList []DescribeTrafficControlsTrafficControl

func (list *DescribeTrafficControlsTrafficControlList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsTrafficControl)
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

type DescribeTrafficControlsSpecialPolicyList []DescribeTrafficControlsSpecialPolicy

func (list *DescribeTrafficControlsSpecialPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsSpecialPolicy)
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

type DescribeTrafficControlsSpecialList []DescribeTrafficControlsSpecial

func (list *DescribeTrafficControlsSpecialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsSpecial)
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
