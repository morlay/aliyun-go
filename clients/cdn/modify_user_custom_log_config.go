package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserCustomLogConfigRequest struct {
	requests.RpcRequest
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
	Tag           string `position:"Query" name:"Tag"`
}

func (req *ModifyUserCustomLogConfigRequest) Invoke(client *sdk.Client) (resp *ModifyUserCustomLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyUserCustomLogConfig", "", "")
	resp = &ModifyUserCustomLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserCustomLogConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
