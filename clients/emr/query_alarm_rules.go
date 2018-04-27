package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAlarmRulesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *QueryAlarmRulesRequest) Invoke(client *sdk.Client) (resp *QueryAlarmRulesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryAlarmRules", "", "")
	resp = &QueryAlarmRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAlarmRulesResponse struct {
	responses.BaseResponse
	RequestId string
	AlarmList QueryAlarmRulesAlarmList
}

type QueryAlarmRulesAlarm struct {
	Name               string
	MetricName         string
	Period             int
	Threshold          string
	EvaluationCount    int
	StartTime          int
	EndTime            int
	SilenceTime        int
	NotifyType         int
	Enable             bool
	State              string
	ComparisonOperator string
	ContactGroups      string
}

type QueryAlarmRulesAlarmList []QueryAlarmRulesAlarm

func (list *QueryAlarmRulesAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAlarmRulesAlarm)
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
