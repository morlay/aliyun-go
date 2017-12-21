package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetPageCompressConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetPageCompressConfigRequest) Invoke(client *sdk.Client) (response *SetPageCompressConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetPageCompressConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetPageCompressConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetPageCompressConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetPageCompressConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetPageCompressConfigResponse struct {
	RequestId string
}
