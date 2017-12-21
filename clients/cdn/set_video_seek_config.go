package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetVideoSeekConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetVideoSeekConfigRequest) Invoke(client *sdk.Client) (response *SetVideoSeekConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetVideoSeekConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetVideoSeekConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetVideoSeekConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetVideoSeekConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetVideoSeekConfigResponse struct {
	RequestId string
}
