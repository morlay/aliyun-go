package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code       common.String
	Message    common.String
	RequestId  common.String
	Datapoints ListActiveAlertRuleAlarmList
}

type ListActiveAlertRuleAlarm struct {
	Uuid               common.String
	Name               common.String
	Namespace          common.String
	MetricName         common.String
	Period             common.String
	Statistics         common.String
	ComparisonOperator common.String
	Threshold          common.String
	EvaluationCount    common.String
	StartTime          common.String
	EndTime            common.String
	SilenceTime        common.String
	NotifyType         common.String
	Enable             common.String
	State              common.String
	ContactGroups      common.String
	Webhook            common.String
	RuleName           common.String
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
