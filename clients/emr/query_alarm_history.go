package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryAlarmHistoryRequest struct {
	requests.RpcRequest
	Cursor          string `position:"Query" name:"Cursor"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Size            int    `position:"Query" name:"Size"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StartTimeStamp  int64  `position:"Query" name:"StartTimeStamp"`
	EndTimeStamp    int64  `position:"Query" name:"EndTimeStamp"`
}

func (req *QueryAlarmHistoryRequest) Invoke(client *sdk.Client) (resp *QueryAlarmHistoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryAlarmHistory", "", "")
	resp = &QueryAlarmHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAlarmHistoryResponse struct {
	responses.BaseResponse
	RequestId        common.String
	Total            common.String
	Cursor           common.String
	AlarmHistoryList QueryAlarmHistoryEmrAlarmHistoryList
}

type QueryAlarmHistoryEmrAlarmHistory struct {
	ClusterId     common.String
	Role          common.String
	InstanceId    common.String
	Name          common.String
	MetricName    common.String
	AlarmTime     common.Long
	LastTime      common.Long
	State         common.String
	Status        common.Integer
	ContactGroups common.String
}

type QueryAlarmHistoryEmrAlarmHistoryList []QueryAlarmHistoryEmrAlarmHistory

func (list *QueryAlarmHistoryEmrAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAlarmHistoryEmrAlarmHistory)
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
