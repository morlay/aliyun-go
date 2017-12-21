package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DeleteCasterRequest) Invoke(client *sdk.Client) (response *DeleteCasterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCasterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCaster", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCasterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCasterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCasterResponse struct {
	RequestId string
	CasterId  string
}
