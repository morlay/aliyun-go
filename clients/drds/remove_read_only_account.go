package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveReadOnlyAccountRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r RemoveReadOnlyAccountRequest) Invoke(client *sdk.Client) (response *RemoveReadOnlyAccountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveReadOnlyAccountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "RemoveReadOnlyAccount", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveReadOnlyAccountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveReadOnlyAccountResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveReadOnlyAccountResponse struct {
	RequestId string
	Success   bool
}
