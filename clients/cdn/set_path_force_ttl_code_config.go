package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetPathForceTtlCodeConfigRequest struct {
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	CodeString    string `position:"Query" name:"CodeString"`
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

func (r SetPathForceTtlCodeConfigRequest) Invoke(client *sdk.Client) (response *SetPathForceTtlCodeConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetPathForceTtlCodeConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetPathForceTtlCodeConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetPathForceTtlCodeConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetPathForceTtlCodeConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetPathForceTtlCodeConfigResponse struct {
	RequestId string
}
