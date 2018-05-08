package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	TotalCount       common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	AlarmHistoryList ListEmrAlarmForAdminAlarmHistoryList
}

type ListEmrAlarmForAdminAlarmHistory struct {
	Id           common.Long
	EventType    common.String
	Topic        common.String
	UniqueKey    common.String
	EntityId     common.String
	Priority     common.Integer
	Body         common.String
	Status       common.String
	ClusterBizId common.String
	GmtCreate    common.String
	GmtModified  common.String
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
