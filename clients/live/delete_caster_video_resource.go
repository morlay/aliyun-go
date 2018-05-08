package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasterVideoResourceRequest struct {
	requests.RpcRequest
	ResourceId string `position:"Query" name:"ResourceId"`
	CasterId   string `position:"Query" name:"CasterId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *DeleteCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterVideoResource", "live", "")
	resp = &DeleteCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
