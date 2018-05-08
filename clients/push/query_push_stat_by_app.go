package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	AppPushStats QueryPushStatByAppAppPushStatList
}

type QueryPushStatByAppAppPushStat struct {
	Time                   common.String
	AcceptCount            common.Long
	SentCount              common.Long
	ReceivedCount          common.Long
	OpenedCount            common.Long
	DeletedCount           common.Long
	SmsSentCount           common.Long
	SmsSkipCount           common.Long
	SmsFailedCount         common.Long
	SmsReceiveSuccessCount common.Long
	SmsReceiveFailedCount  common.Long
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
