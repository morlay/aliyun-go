package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCasterComponentRequest struct {
	requests.RpcRequest
	ComponentType       string `position:"Query" name:"ComponentType"`
	LocationId          string `position:"Query" name:"LocationId"`
	ImageLayerContent   string `position:"Query" name:"ImageLayerContent"`
	CasterId            string `position:"Query" name:"CasterId"`
	Effect              string `position:"Query" name:"Effect"`
	ComponentLayer      string `position:"Query" name:"ComponentLayer"`
	CaptionLayerContent string `position:"Query" name:"CaptionLayerContent"`
	ComponentName       string `position:"Query" name:"ComponentName"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	TextLayerContent    string `position:"Query" name:"TextLayerContent"`
}

func (req *AddCasterComponentRequest) Invoke(client *sdk.Client) (resp *AddCasterComponentResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterComponent", "live", "")
	resp = &AddCasterComponentResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterComponentResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ComponentId common.String
}
