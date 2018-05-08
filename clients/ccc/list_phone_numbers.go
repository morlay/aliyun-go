package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	PhoneNumbers   ListPhoneNumbersPhoneNumberList
}

type ListPhoneNumbersPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
	Province               common.String
	City                   common.String
	ContactFlow            ListPhoneNumbersContactFlow
}

type ListPhoneNumbersContactFlow struct {
	ContactFlowId          common.String
	InstanceId             common.String
	ContactFlowName        common.String
	ContactFlowDescription common.String
	Type                   common.String
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
