package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPushListRequest struct {
	requests.RpcRequest
	PageSize  int    `position:"Query" name:"PageSize"`
	EndTime   string `position:"Query" name:"EndTime"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	StartTime string `position:"Query" name:"StartTime"`
	Page      int    `position:"Query" name:"Page"`
	PushType  string `position:"Query" name:"PushType"`
}

func (req *QueryPushListRequest) Invoke(client *sdk.Client) (resp *QueryPushListResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryPushList", "", "")
	resp = &QueryPushListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPushListResponse struct {
	responses.BaseResponse
	RequestId        string
	HasNext          bool
	Page             int
	PageSize         int
	PushMessageInfos QueryPushListPushMessageInfoList
}

type QueryPushListPushMessageInfo struct {
	AppKey     int64
	AppName    string
	MessageId  string
	Type       string
	DeviceType string
	PushTime   string
	Title      string
	Body       string
}

type QueryPushListPushMessageInfoList []QueryPushListPushMessageInfo

func (list *QueryPushListPushMessageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushListPushMessageInfo)
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
