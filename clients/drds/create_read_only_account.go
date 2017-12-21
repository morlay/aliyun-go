package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateReadOnlyAccountRequest struct {
	requests.RpcRequest
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *CreateReadOnlyAccountRequest) Invoke(client *sdk.Client) (resp *CreateReadOnlyAccountResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateReadOnlyAccount", "", "")
	resp = &CreateReadOnlyAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateReadOnlyAccountResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Data      CreateReadOnlyAccountData
}

type CreateReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
