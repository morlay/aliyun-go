package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPushRecordsRequest struct {
	requests.RpcRequest
	PageSize  int    `position:"Query" name:"PageSize"`
	EndTime   string `position:"Query" name:"EndTime"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	StartTime string `position:"Query" name:"StartTime"`
	Page      int    `position:"Query" name:"Page"`
	PushType  string `position:"Query" name:"PushType"`
}

func (req *ListPushRecordsRequest) Invoke(client *sdk.Client) (resp *ListPushRecordsResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "ListPushRecords", "", "")
	resp = &ListPushRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPushRecordsResponse struct {
	responses.BaseResponse
	RequestId        string
	Total            int
	Page             int
	PageSize         int
	PushMessageInfos ListPushRecordsPushMessageInfoList
}

type ListPushRecordsPushMessageInfo struct {
	AppKey     int64
	AppName    string
	MessageId  string
	Type       string
	DeviceType string
	PushTime   string
	Title      string
	Body       string
}

type ListPushRecordsPushMessageInfoList []ListPushRecordsPushMessageInfo

func (list *ListPushRecordsPushMessageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPushRecordsPushMessageInfo)
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
