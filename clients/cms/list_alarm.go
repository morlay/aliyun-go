package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAlarmRequest struct {
	requests.RpcRequest
	IsEnable         string `position:"Query" name:"IsEnable"`
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Name             string `position:"Query" name:"Name"`
	Namespace        string `position:"Query" name:"Namespace"`
	PageSize         int    `position:"Query" name:"PageSize"`
	Id               string `position:"Query" name:"Id"`
	State            string `position:"Query" name:"State"`
	Dimension        string `position:"Query" name:"Dimension"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
}

func (req *ListAlarmRequest) Invoke(client *sdk.Client) (resp *ListAlarmResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListAlarm", "cms", "")
	resp = &ListAlarmResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAlarmResponse struct {
	responses.BaseResponse
	Success   bool
	Code      common.String
	Message   common.String
	NextToken common.Integer
	Total     common.Integer
	RequestId common.String
	AlarmList ListAlarmAlarmList
}

type ListAlarmAlarm struct {
	Id                 common.String
	Name               common.String
	Namespace          common.String
	MetricName         common.String
	Dimensions         common.String
	Period             common.Integer
	Statistics         common.String
	ComparisonOperator common.String
	Threshold          common.String
	EvaluationCount    common.Integer
	StartTime          common.Integer
	EndTime            common.Integer
	SilenceTime        common.Integer
	NotifyType         common.Integer
	Enable             bool
	State              common.String
	ContactGroups      common.String
	Webhook            common.String
}

type ListAlarmAlarmList []ListAlarmAlarm

func (list *ListAlarmAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmAlarm)
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
