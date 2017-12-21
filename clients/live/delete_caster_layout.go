package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterLayoutRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	LayoutId      string `position:"Query" name:"LayoutId"`
}

func (r DeleteCasterLayoutRequest) Invoke(client *sdk.Client) (response *DeleteCasterLayoutResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCasterLayoutRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterLayout", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCasterLayoutResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCasterLayoutResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCasterLayoutResponse struct {
	RequestId string
	CasterId  string
	LayoutId  string
}
