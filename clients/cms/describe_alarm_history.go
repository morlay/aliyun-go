package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAlarmHistoryRequest struct {
	AlertName  string `position:"Query" name:"AlertName"`
	GroupId    string `position:"Query" name:"GroupId"`
	EndTime    string `position:"Query" name:"EndTime"`
	RuleName   string `position:"Query" name:"RuleName"`
	StartTime  string `position:"Query" name:"StartTime"`
	Ascending  string `position:"Query" name:"Ascending"`
	OnlyCount  string `position:"Query" name:"OnlyCount"`
	Namespace  string `position:"Query" name:"Namespace"`
	PageSize   int    `position:"Query" name:"PageSize"`
	State      string `position:"Query" name:"State"`
	Page       int    `position:"Query" name:"Page"`
	MetricName string `position:"Query" name:"MetricName"`
	Status     string `position:"Query" name:"Status"`
}

func (r DescribeAlarmHistoryRequest) Invoke(client *sdk.Client) (response *DescribeAlarmHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAlarmHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "DescribeAlarmHistory", "cms", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAlarmHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAlarmHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAlarmHistoryResponse struct {
	Success          bool
	Code             string
	Message          string
	Total            string
	RequestId        string
	AlarmHistoryList DescribeAlarmHistoryAlarmHistoryList
}

type DescribeAlarmHistoryAlarmHistory struct {
	Id              string
	AlertName       string
	GroupId         string
	Namespace       string
	MetricName      string
	Dimensions      string
	Expression      string
	EvaluationCount int
	Value           string
	AlertTime       int64
	LastTime        int64
	Level           string
	PreLevel        string
	RuleName        string
	State           string
	Status          int
	UserId          string
	Webhooks        string
	ContactGroups   DescribeAlarmHistoryContactGroupList
	Contacts        DescribeAlarmHistoryContactList
	ContactALIIMs   DescribeAlarmHistoryContactALIIMList
	ContactSmses    DescribeAlarmHistoryContactSmseList
	ContactMails    DescribeAlarmHistoryContactMailList
}

type DescribeAlarmHistoryAlarmHistoryList []DescribeAlarmHistoryAlarmHistory

func (list *DescribeAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAlarmHistoryAlarmHistory)
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

type DescribeAlarmHistoryContactGroupList []string

func (list *DescribeAlarmHistoryContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeAlarmHistoryContactList []string

func (list *DescribeAlarmHistoryContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeAlarmHistoryContactALIIMList []string

func (list *DescribeAlarmHistoryContactALIIMList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeAlarmHistoryContactSmseList []string

func (list *DescribeAlarmHistoryContactSmseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeAlarmHistoryContactMailList []string

func (list *DescribeAlarmHistoryContactMailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
