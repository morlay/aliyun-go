package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadRuleRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadRuleRequest) Invoke(client *sdk.Client) (response *UploadRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadRule", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadRuleResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      UploadRuleDatumList
}

type UploadRuleDatumList []string

func (list *UploadRuleDatumList) UnmarshalJSON(data []byte) error {
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
