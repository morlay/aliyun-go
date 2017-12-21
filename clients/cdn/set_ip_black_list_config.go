package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetIpBlackListConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
}

func (r SetIpBlackListConfigRequest) Invoke(client *sdk.Client) (response *SetIpBlackListConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetIpBlackListConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetIpBlackListConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetIpBlackListConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetIpBlackListConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetIpBlackListConfigResponse struct {
	RequestId string
}
