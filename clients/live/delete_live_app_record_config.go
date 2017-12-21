package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveAppRecordConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveAppRecordConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveAppRecordConfig", "", "")
	resp = &DeleteLiveAppRecordConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveAppRecordConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
