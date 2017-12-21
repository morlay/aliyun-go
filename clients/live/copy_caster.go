package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CopyCasterRequest struct {
	SrcCasterId   string `position:"Query" name:"SrcCasterId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterName    string `position:"Query" name:"CasterName"`
	ClientToken   string `position:"Query" name:"ClientToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r CopyCasterRequest) Invoke(client *sdk.Client) (response *CopyCasterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CopyCasterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "CopyCaster", "", "")

	resp := struct {
		*responses.BaseResponse
		CopyCasterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CopyCasterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CopyCasterResponse struct {
	RequestId string
	CasterId  string
}
