package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListContactFlowsRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListContactFlowsRequest) Invoke(client *sdk.Client) (resp *ListContactFlowsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListContactFlows", "CCC", "")
	resp = &ListContactFlowsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListContactFlowsResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	ContactFlows   ListContactFlowsContactFlowList
}

type ListContactFlowsContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
	AppliedVersion         string
	Versions               ListContactFlowsContactFlowVersionList
	PhoneNumbers           ListContactFlowsPhoneNumberList
}

type ListContactFlowsContactFlowVersion struct {
	ContactFlowVersionId          string
	Version                       string
	ContactFlowVersionDescription string
	LastModified                  string
	LastModifiedBy                string
	LockedBy                      string
	Status                        string
}

type ListContactFlowsPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               bool
	RemainingTime          int
	AllowOutbound          bool
	Usage                  string
	Trunks                 int
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
