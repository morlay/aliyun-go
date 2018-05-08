package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasterComponentRequest struct {
	requests.RpcRequest
	ComponentId string `position:"Query" name:"ComponentId"`
	CasterId    string `position:"Query" name:"CasterId"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCasterComponentRequest) Invoke(client *sdk.Client) (resp *DeleteCasterComponentResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterComponent", "live", "")
	resp = &DeleteCasterComponentResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterComponentResponse struct {
	responses.BaseResponse
	RequestId   common.String
	CasterId    common.String
	ComponentId common.String
}
