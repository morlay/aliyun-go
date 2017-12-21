package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterComponentRequest struct {
	ComponentId   string `position:"Query" name:"ComponentId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DeleteCasterComponentRequest) Invoke(client *sdk.Client) (response *DeleteCasterComponentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCasterComponentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterComponent", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCasterComponentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCasterComponentResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCasterComponentResponse struct {
	RequestId   string
	CasterId    string
	ComponentId string
}
