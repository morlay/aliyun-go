package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterVideoResourceRequest struct {
	requests.RpcRequest
	ResourceId    string `position:"Query" name:"ResourceId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DeleteCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *DeleteCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterVideoResource", "", "")
	resp = &DeleteCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId string
}
