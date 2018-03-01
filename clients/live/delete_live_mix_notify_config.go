package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveMixNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveMixNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveMixNotifyConfig", "live", "")
	resp = &DeleteLiveMixNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveMixNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
