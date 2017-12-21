package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateEnterpriseRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	BoxNumber   int    `position:"Query" name:"BoxNumber"`
	ServiceFlag string `position:"Query" name:"ServiceFlag"`
}

func (r CreateEnterpriseRequest) Invoke(client *sdk.Client) (response *CreateEnterpriseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateEnterpriseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "CreateEnterprise", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateEnterpriseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateEnterpriseResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateEnterpriseResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList CreateEnterpriseErrorMessageList
}

type CreateEnterpriseErrorMessage struct {
	ErrorMessage string
}

type CreateEnterpriseErrorMessageList []CreateEnterpriseErrorMessage

func (list *CreateEnterpriseErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateEnterpriseErrorMessage)
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
