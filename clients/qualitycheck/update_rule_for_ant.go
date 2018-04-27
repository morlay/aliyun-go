package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRuleForAntRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UpdateRuleForAntRequest) Invoke(client *sdk.Client) (resp *UpdateRuleForAntResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UpdateRuleForAnt", "", "")
	resp = &UpdateRuleForAntResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateRuleForAntResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      UpdateRuleForAntDatumList
}

type UpdateRuleForAntDatumList []string

func (list *UpdateRuleForAntDatumList) UnmarshalJSON(data []byte) error {
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
