package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveAppRecordConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveAppRecordConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveAppRecordConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveAppRecordConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveAppRecordConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveAppRecordConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveAppRecordConfigResponse struct {
	RequestId string
}
