package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLivePullStreamInfoConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DeleteLivePullStreamInfoConfigRequest) Invoke(client *sdk.Client) (response *DeleteLivePullStreamInfoConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLivePullStreamInfoConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLivePullStreamInfoConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLivePullStreamInfoConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLivePullStreamInfoConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLivePullStreamInfoConfigResponse struct {
	RequestId string
}
