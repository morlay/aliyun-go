package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddLiveMixNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveMixNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveMixNotifyConfig", "live", "")
	resp = &AddLiveMixNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveMixNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
