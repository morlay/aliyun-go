package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDrdsDBPasswdRequest struct {
	requests.RpcRequest
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *ModifyDrdsDBPasswdRequest) Invoke(client *sdk.Client) (resp *ModifyDrdsDBPasswdResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsDBPasswd", "", "")
	resp = &ModifyDrdsDBPasswdResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDrdsDBPasswdResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
