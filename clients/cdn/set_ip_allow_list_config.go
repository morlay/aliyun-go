package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetIpAllowListConfigRequest struct {
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetIpAllowListConfigRequest) Invoke(client *sdk.Client) (response *SetIpAllowListConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetIpAllowListConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetIpAllowListConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetIpAllowListConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetIpAllowListConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetIpAllowListConfigResponse struct {
	RequestId string
}
