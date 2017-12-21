package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveRecordNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveRecordNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveRecordNotifyConfig", "", "")
	resp = &DeleteLiveRecordNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveRecordNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
