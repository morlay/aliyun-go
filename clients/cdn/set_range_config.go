package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetRangeConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetRangeConfigRequest) Invoke(client *sdk.Client) (response *SetRangeConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetRangeConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRangeConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetRangeConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetRangeConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetRangeConfigResponse struct {
	RequestId string
}
