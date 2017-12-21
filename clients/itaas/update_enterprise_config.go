package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateEnterpriseConfigRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	ConfigKey   string `position:"Query" name:"ConfigKey"`
	ConfigValue string `position:"Query" name:"ConfigValue"`
	Memo        string `position:"Query" name:"Memo"`
}

func (req *UpdateEnterpriseConfigRequest) Invoke(client *sdk.Client) (resp *UpdateEnterpriseConfigResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "UpdateEnterpriseConfig", "", "")
	resp = &UpdateEnterpriseConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateEnterpriseConfigResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateEnterpriseConfigErrorMessageList
}

type UpdateEnterpriseConfigErrorMessage struct {
	ErrorMessage string
}

type UpdateEnterpriseConfigErrorMessageList []UpdateEnterpriseConfigErrorMessage

func (list *UpdateEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateEnterpriseConfigErrorMessage)
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
