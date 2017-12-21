package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetReqHeaderConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Value         string `position:"Query" name:"Value"`
	Key           string `position:"Query" name:"Key"`
}

func (r SetReqHeaderConfigRequest) Invoke(client *sdk.Client) (response *SetReqHeaderConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetReqHeaderConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetReqHeaderConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetReqHeaderConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetReqHeaderConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetReqHeaderConfigResponse struct {
	RequestId string
}
