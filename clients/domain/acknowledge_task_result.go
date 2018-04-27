package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AcknowledgeTaskResultRequest struct {
	requests.RpcRequest
	TaskDetailNos *AcknowledgeTaskResultTaskDetailNoList `position:"Query" type:"Repeated" name:"TaskDetailNo"`
	UserClientIp  string                                 `position:"Query" name:"UserClientIp"`
	Lang          string                                 `position:"Query" name:"Lang"`
}

func (req *AcknowledgeTaskResultRequest) Invoke(client *sdk.Client) (resp *AcknowledgeTaskResultResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "AcknowledgeTaskResult", "", "")
	resp = &AcknowledgeTaskResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type AcknowledgeTaskResultResponse struct {
	responses.BaseResponse
	RequestId string
	Result    int
}

type AcknowledgeTaskResultTaskDetailNoList []string

func (list *AcknowledgeTaskResultTaskDetailNoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
