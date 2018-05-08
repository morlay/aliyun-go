package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListContactFlowsRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListContactFlowsRequest) Invoke(client *sdk.Client) (resp *ListContactFlowsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListContactFlows", "ccc", "")
	resp = &ListContactFlowsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListContactFlowsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	ContactFlows   ListContactFlowsContactFlowList
}

type ListContactFlowsContactFlow struct {
	ContactFlowId          common.String
	InstanceId             common.String
	ContactFlowName        common.String
	ContactFlowDescription common.String
	Type                   common.String
	AppliedVersion         common.String
	Versions               ListContactFlowsContactFlowVersionList
	PhoneNumbers           ListContactFlowsPhoneNumberList
}

type ListContactFlowsContactFlowVersion struct {
	ContactFlowVersionId          common.String
	Version                       common.String
	ContactFlowVersionDescription common.String
	LastModified                  common.String
	LastModifiedBy                common.String
	LockedBy                      common.String
	Status                        common.String
}

type ListContactFlowsPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
}

type ListContactFlowsContactFlowList []ListContactFlowsContactFlow

func (list *ListContactFlowsContactFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlow)
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

type ListContactFlowsContactFlowVersionList []ListContactFlowsContactFlowVersion

func (list *ListContactFlowsContactFlowVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlowVersion)
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

type ListContactFlowsPhoneNumberList []ListContactFlowsPhoneNumber

func (list *ListContactFlowsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsPhoneNumber)
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
