package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAlarmHistoryRequest struct {
	requests.RpcRequest
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

func (req *DescribeAlarmHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeAlarmHistoryResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DescribeAlarmHistory", "cms", "")
	resp = &DescribeAlarmHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAlarmHistoryResponse struct {
	responses.BaseResponse
	Success          bool
	Code             common.String
	Message          common.String
	Total            common.String
	RequestId        common.String
	AlarmHistoryList DescribeAlarmHistoryAlarmHistoryList
}

type DescribeAlarmHistoryAlarmHistory struct {
	Id              common.String
	AlertName       common.String
	GroupId         common.String
	Namespace       common.String
	MetricName      common.String
	Dimensions      common.String
	Expression      common.String
	EvaluationCount common.Integer
	Value           common.String
	AlertTime       common.Long
	LastTime        common.Long
	Level           common.String
	PreLevel        common.String
	RuleName        common.String
	State           common.String
	Status          common.Integer
	UserId          common.String
	Webhooks        common.String
	InstanceName    common.String
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

type DescribeAlarmHistoryContactGroupList []common.String

func (list *DescribeAlarmHistoryContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeAlarmHistoryContactList []common.String

func (list *DescribeAlarmHistoryContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeAlarmHistoryContactALIIMList []common.String

func (list *DescribeAlarmHistoryContactALIIMList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeAlarmHistoryContactSmseList []common.String

func (list *DescribeAlarmHistoryContactSmseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeAlarmHistoryContactMailList []common.String

func (list *DescribeAlarmHistoryContactMailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
