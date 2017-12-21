package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTrafficControlsRequest struct {
	TrafficControlId   string `position:"Query" name:"TrafficControlId"`
	GroupId            string `position:"Query" name:"GroupId"`
	ApiId              string `position:"Query" name:"ApiId"`
	StageName          string `position:"Query" name:"StageName"`
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	PageNumber         int    `position:"Query" name:"PageNumber"`
	PageSize           int    `position:"Query" name:"PageSize"`
}

func (r DescribeTrafficControlsRequest) Invoke(client *sdk.Client) (response *DescribeTrafficControlsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTrafficControlsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeTrafficControls", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTrafficControlsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTrafficControlsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTrafficControlsResponse struct {
	RequestId       string
	TotalCount      int
	PageSize        int
	PageNumber      int
	TrafficControls DescribeTrafficControlsTrafficControlList
}

type DescribeTrafficControlsTrafficControl struct {
	TrafficControlId   string
	TrafficControlName string
	Description        string
	TrafficControlUnit string
	ApiDefault         int
	UserDefault        int
	AppDefault         int
	CreatedTime        string
	ModifiedTime       string
	SpecialPolicies    DescribeTrafficControlsSpecialPolicyList
}

type DescribeTrafficControlsSpecialPolicy struct {
	SpecialType string
	Specials    DescribeTrafficControlsSpecialList
}

type DescribeTrafficControlsSpecial struct {
	SpecialKey   string
	TrafficValue int
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
