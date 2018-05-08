package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	PushStats QueryPushStatByMsgPushStatList
}

type QueryPushStatByMsgPushStat struct {
	MessageId              common.String
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
