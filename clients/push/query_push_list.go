package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	HasNext          bool
	Page             common.Integer
	PageSize         common.Integer
	PushMessageInfos QueryPushListPushMessageInfoList
}

type QueryPushListPushMessageInfo struct {
	AppKey     common.Long
	AppName    common.String
	MessageId  common.String
	Type       common.String
	DeviceType common.String
	PushTime   common.String
	Title      common.String
	Body       common.String
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
