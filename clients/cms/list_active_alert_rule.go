package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListActiveAlertRuleRequest struct {
	requests.RpcRequest
	Product string `position:"Query" name:"Product"`
	UserId  string `position:"Query" name:"UserId"`
}

func (req *ListActiveAlertRuleRequest) Invoke(client *sdk.Client) (resp *ListActiveAlertRuleResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListActiveAlertRule", "cms", "")
	resp = &ListActiveAlertRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListActiveAlertRuleResponse struct {
	responses.BaseResponse
	Success    bool
	Code       string
	Message    string
	RequestId  string
	Datapoints ListActiveAlertRuleAlarmList
}

type ListActiveAlertRuleAlarm struct {
	Uuid               string
	Name               string
	Namespace          string
	MetricName         string
	Period             string
	Statistics         string
	ComparisonOperator string
	Threshold          string
	EvaluationCount    string
	StartTime          string
	EndTime            string
	SilenceTime        string
	NotifyType         string
	Enable             string
	State              string
	ContactGroups      string
	Webhook            string
	RuleName           string
}

type ListActiveAlertRuleAlarmList []ListActiveAlertRuleAlarm

func (list *ListActiveAlertRuleAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListActiveAlertRuleAlarm)
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
