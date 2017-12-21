package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDynamicConfigRequest struct {
	DynamicOrigin       string `position:"Query" name:"DynamicOrigin"`
	StaticType          string `position:"Query" name:"StaticType"`
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	StaticUri           string `position:"Query" name:"StaticUri"`
	DomainName          string `position:"Query" name:"DomainName"`
	StaticPath          string `position:"Query" name:"StaticPath"`
	DynamicCacheControl string `position:"Query" name:"DynamicCacheControl"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
}

func (r SetDynamicConfigRequest) Invoke(client *sdk.Client) (response *SetDynamicConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDynamicConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetDynamicConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDynamicConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDynamicConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDynamicConfigResponse struct {
	RequestId string
}
