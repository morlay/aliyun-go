package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhoneNumbersRequest struct {
	requests.RpcRequest
	OutboundOnly string `position:"Query" name:"OutboundOnly"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

func (req *ListPhoneNumbersRequest) Invoke(client *sdk.Client) (resp *ListPhoneNumbersResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListPhoneNumbers", "ccc", "")
	resp = &ListPhoneNumbersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhoneNumbersResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumbers   ListPhoneNumbersPhoneNumberList
}

type ListPhoneNumbersPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               bool
	RemainingTime          int
	AllowOutbound          bool
	Usage                  string
	Trunks                 int
	Province               string
	City                   string
	ContactFlow            ListPhoneNumbersContactFlow
}

type ListPhoneNumbersContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}

type ListPhoneNumbersPhoneNumberList []ListPhoneNumbersPhoneNumber

func (list *ListPhoneNumbersPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhoneNumbersPhoneNumber)
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
