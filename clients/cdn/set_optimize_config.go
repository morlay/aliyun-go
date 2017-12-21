package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetOptimizeConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetOptimizeConfigRequest) Invoke(client *sdk.Client) (response *SetOptimizeConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetOptimizeConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetOptimizeConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetOptimizeConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetOptimizeConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetOptimizeConfigResponse struct {
	RequestId string
}
