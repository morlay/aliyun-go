package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyReadOnlyAccountPasswordRequest struct {
	requests.RpcRequest
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	OriginPassword string `position:"Query" name:"OriginPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *ModifyReadOnlyAccountPasswordRequest) Invoke(client *sdk.Client) (resp *ModifyReadOnlyAccountPasswordResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyReadOnlyAccountPassword", "", "")
	resp = &ModifyReadOnlyAccountPasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReadOnlyAccountPasswordResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
