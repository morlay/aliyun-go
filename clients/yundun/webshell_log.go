package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WebshellLogRequest struct {
	requests.RpcRequest
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int    `position:"Query" name:"RecordType"`
}

func (req *WebshellLogRequest) Invoke(client *sdk.Client) (resp *WebshellLogResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "WebshellLog", "", "")
	resp = &WebshellLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type WebshellLogResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    WebshellLogWebshellLogList
}

type WebshellLogWebshellLog struct {
	Id     string
	Path   string
	Status int
	Time   string
}

type WebshellLogWebshellLogList []WebshellLogWebshellLog

func (list *WebshellLogWebshellLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WebshellLogWebshellLog)
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
