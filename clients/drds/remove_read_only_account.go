package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveReadOnlyAccountRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *RemoveReadOnlyAccountRequest) Invoke(client *sdk.Client) (resp *RemoveReadOnlyAccountResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "RemoveReadOnlyAccount", "", "")
	resp = &RemoveReadOnlyAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveReadOnlyAccountResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
