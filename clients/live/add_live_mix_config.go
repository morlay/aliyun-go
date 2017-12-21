package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveMixConfigRequest struct {
	requests.RpcRequest
	Template      string `position:"Query" name:"Template"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveMixConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveMixConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveMixConfig", "", "")
	resp = &AddLiveMixConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveMixConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
