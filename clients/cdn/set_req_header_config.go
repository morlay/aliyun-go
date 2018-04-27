package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetReqHeaderConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      int64  `position:"Query" name:"ConfigId"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Value         string `position:"Query" name:"Value"`
	Key           string `position:"Query" name:"Key"`
}

func (req *SetReqHeaderConfigRequest) Invoke(client *sdk.Client) (resp *SetReqHeaderConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetReqHeaderConfig", "", "")
	resp = &SetReqHeaderConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetReqHeaderConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
