package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateReadOnlyAccountRequest struct {
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r CreateReadOnlyAccountRequest) Invoke(client *sdk.Client) (response *CreateReadOnlyAccountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateReadOnlyAccountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateReadOnlyAccount", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateReadOnlyAccountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateReadOnlyAccountResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateReadOnlyAccountResponse struct {
	RequestId string
	Success   bool
	Data      CreateReadOnlyAccountData
}

type CreateReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
