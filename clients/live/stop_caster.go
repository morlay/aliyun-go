package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCasterRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r StopCasterRequest) Invoke(client *sdk.Client) (response *StopCasterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopCasterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "StopCaster", "", "")

	resp := struct {
		*responses.BaseResponse
		StopCasterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopCasterResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopCasterResponse struct {
	RequestId string
}
