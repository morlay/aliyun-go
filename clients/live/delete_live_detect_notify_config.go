package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveDetectNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveDetectNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveDetectNotifyConfig", "", "")
	resp = &DeleteLiveDetectNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveDetectNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
