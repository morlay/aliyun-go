package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveMixNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveMixNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveMixNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveMixNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveMixNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveMixNotifyConfigResponse struct {
	RequestId string
}
