package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveMixConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveMixConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveMixConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveMixConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveMixConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveMixConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveMixConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveMixConfigResponse struct {
	RequestId string
}
