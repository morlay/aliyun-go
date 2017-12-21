package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveRecordNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveRecordNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveRecordNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveRecordNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveRecordNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveRecordNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveRecordNotifyConfigResponse struct {
	RequestId string
}
