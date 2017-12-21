package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetOptimizeConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetOptimizeConfigRequest) Invoke(client *sdk.Client) (resp *SetOptimizeConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetOptimizeConfig", "", "")
	resp = &SetOptimizeConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetOptimizeConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
