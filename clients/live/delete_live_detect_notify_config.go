package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveDetectNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveDetectNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveDetectNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveDetectNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveDetectNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveDetectNotifyConfigResponse struct {
	RequestId string
}
