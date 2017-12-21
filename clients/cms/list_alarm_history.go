package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAlarmHistoryRequest struct {
	Cursor           string `position:"Query" name:"Cursor"`
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Size             int    `position:"Query" name:"Size"`
	EndTime          string `position:"Query" name:"EndTime"`
	Id               string `position:"Query" name:"Id"`
	StartTime        string `position:"Query" name:"StartTime"`
}

func (r ListAlarmHistoryRequest) Invoke(client *sdk.Client) (response *ListAlarmHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAlarmHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "ListAlarmHistory", "cms", "")

	resp := struct {
		*responses.BaseResponse
		ListAlarmHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAlarmHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAlarmHistoryResponse struct {
	Success          bool
	Code             string
	Message          string
	Cursor           string
	RequestId        string
	AlarmHistoryList ListAlarmHistoryAlarmHistoryList
}

type ListAlarmHistoryAlarmHistory struct {
	Id              string
	Name            string
	Namespace       string
	MetricName      string
	Dimension       string
	EvaluationCount int
	Value           string
	AlarmTime       int64
	LastTime        int64
	State           string
	Status          int
	ContactGroups   string
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
