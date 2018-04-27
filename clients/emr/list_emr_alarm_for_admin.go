package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListEmrAlarmForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	UniqueKey       string `position:"Query" name:"UniqueKey"`
	PageSize        int    `position:"Query" name:"PageSize"`
	Topic           string `position:"Query" name:"Topic"`
	EventType       string `position:"Query" name:"EventType"`
	EntityId        string `position:"Query" name:"EntityId"`
	Priority        int    `position:"Query" name:"Priority"`
	ClusterBizId    string `position:"Query" name:"ClusterBizId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListEmrAlarmForAdminRequest) Invoke(client *sdk.Client) (resp *ListEmrAlarmForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListEmrAlarmForAdmin", "", "")
	resp = &ListEmrAlarmForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEmrAlarmForAdminResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	AlarmHistoryList ListEmrAlarmForAdminAlarmHistoryList
}

type ListEmrAlarmForAdminAlarmHistory struct {
	Id           int64
	EventType    string
	Topic        string
	UniqueKey    string
	EntityId     string
	Priority     int
	Body         string
	Status       string
	ClusterBizId string
	GmtCreate    string
	GmtModified  string
}

type ListEmrAlarmForAdminAlarmHistoryList []ListEmrAlarmForAdminAlarmHistory

func (list *ListEmrAlarmForAdminAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmrAlarmForAdminAlarmHistory)
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
