package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      UploadRuleForAntDatumList
}

type UploadRuleForAntDatumList []common.String

func (list *UploadRuleForAntDatumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
