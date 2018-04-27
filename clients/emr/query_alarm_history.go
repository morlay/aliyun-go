package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId        string
	Total            string
	Cursor           string
	AlarmHistoryList QueryAlarmHistoryEmrAlarmHistoryList
}

type QueryAlarmHistoryEmrAlarmHistory struct {
	ClusterId     string
	Role          string
	InstanceId    string
	Name          string
	MetricName    string
	AlarmTime     int64
	LastTime      int64
	State         string
	Status        int
	ContactGroups string
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
