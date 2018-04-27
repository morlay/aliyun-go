package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPushStatByMsgRequest struct {
	requests.RpcRequest
	MessageId int64 `position:"Query" name:"MessageId"`
	AppKey    int64 `position:"Query" name:"AppKey"`
}

func (req *QueryPushStatByMsgRequest) Invoke(client *sdk.Client) (resp *QueryPushStatByMsgResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryPushStatByMsg", "", "")
	resp = &QueryPushStatByMsgResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPushStatByMsgResponse struct {
	responses.BaseResponse
	RequestId string
	PushStats QueryPushStatByMsgPushStatList
}

type QueryPushStatByMsgPushStat struct {
	MessageId              string
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

type QueryPushStatByMsgPushStatList []QueryPushStatByMsgPushStat

func (list *QueryPushStatByMsgPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByMsgPushStat)
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
