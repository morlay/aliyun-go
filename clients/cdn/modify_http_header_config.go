package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyHttpHeaderConfigRequest struct {
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyHttpHeaderConfigRequest) Invoke(client *sdk.Client) (response *ModifyHttpHeaderConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyHttpHeaderConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyHttpHeaderConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyHttpHeaderConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyHttpHeaderConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyHttpHeaderConfigResponse struct {
	RequestId string
}
