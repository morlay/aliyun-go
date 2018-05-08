package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAlarmHistoryRequest struct {
	requests.RpcRequest
	Cursor           string `position:"Query" name:"Cursor"`
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Size             int    `position:"Query" name:"Size"`
	EndTime          string `position:"Query" name:"EndTime"`
	Id               string `position:"Query" name:"Id"`
	StartTime        string `position:"Query" name:"StartTime"`
}

func (req *ListAlarmHistoryRequest) Invoke(client *sdk.Client) (resp *ListAlarmHistoryResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListAlarmHistory", "cms", "")
	resp = &ListAlarmHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAlarmHistoryResponse struct {
	responses.BaseResponse
	Success          bool
	Code             common.String
	Message          common.String
	Cursor           common.String
	RequestId        common.String
	AlarmHistoryList ListAlarmHistoryAlarmHistoryList
}

type ListAlarmHistoryAlarmHistory struct {
	Id              common.String
	Name            common.String
	Namespace       common.String
	MetricName      common.String
	Dimension       common.String
	EvaluationCount common.Integer
	Value           common.String
	AlarmTime       common.Long
	LastTime        common.Long
	State           common.String
	Status          common.Integer
	ContactGroups   common.String
	InstanceName    common.String
}

type ListAlarmHistoryAlarmHistoryList []ListAlarmHistoryAlarmHistory

func (list *ListAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmHistoryAlarmHistory)
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
