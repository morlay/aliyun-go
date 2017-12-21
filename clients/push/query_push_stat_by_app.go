package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPushStatByAppRequest struct {
	requests.RpcRequest
	Granularity string `position:"Query" name:"Granularity"`
	EndTime     string `position:"Query" name:"EndTime"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	StartTime   string `position:"Query" name:"StartTime"`
}

func (req *QueryPushStatByAppRequest) Invoke(client *sdk.Client) (resp *QueryPushStatByAppResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryPushStatByApp", "", "")
	resp = &QueryPushStatByAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPushStatByAppResponse struct {
	responses.BaseResponse
	RequestId    string
	AppPushStats QueryPushStatByAppAppPushStatList
}

type QueryPushStatByAppAppPushStat struct {
	Time                   string
	AcceptCount            int64
	SentCount              int64
	ReceivedCount          int64
	OpenedCount            int64
	DeletedCount           int64
	SmsSentCount           int64
	SmsSkipCount           int64
	SmsFailedCount         int64
	SmsReceiveSuccessCount int64
	SmsReceiveFailedCount  int64
}

type QueryPushStatByAppAppPushStatList []QueryPushStatByAppAppPushStat

func (list *QueryPushStatByAppAppPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByAppAppPushStat)
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
