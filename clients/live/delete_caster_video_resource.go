package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterVideoResourceRequest struct {
	ResourceId    string `position:"Query" name:"ResourceId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DeleteCasterVideoResourceRequest) Invoke(client *sdk.Client) (response *DeleteCasterVideoResourceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCasterVideoResourceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterVideoResource", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCasterVideoResourceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCasterVideoResourceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCasterVideoResourceResponse struct {
	RequestId string
}
