package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadRuleForAntRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadRuleForAntRequest) Invoke(client *sdk.Client) (resp *UploadRuleForAntResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadRuleForAnt", "", "")
	resp = &UploadRuleForAntResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadRuleForAntResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      UploadRuleForAntDatumList
}

type UploadRuleForAntDatumList []string

func (list *UploadRuleForAntDatumList) UnmarshalJSON(data []byte) error {
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
