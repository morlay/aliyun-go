package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	AlarmList QueryAlarmRulesAlarmList
}

type QueryAlarmRulesAlarm struct {
	Name               common.String
	MetricName         common.String
	Period             common.Integer
	Threshold          common.String
	EvaluationCount    common.Integer
	StartTime          common.Integer
	EndTime            common.Integer
	SilenceTime        common.Integer
	NotifyType         common.Integer
	Enable             bool
	State              common.String
	ComparisonOperator common.String
	ContactGroups      common.String
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
