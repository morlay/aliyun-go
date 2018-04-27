package domain_intl

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
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "AcknowledgeTaskResult", "domain", "")
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
