package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterComponentRequest struct {
	requests.RpcRequest
	ComponentId         string `position:"Query" name:"ComponentId"`
	ComponentType       string `position:"Query" name:"ComponentType"`
	ImageLayerContent   string `position:"Query" name:"ImageLayerContent"`
	CasterId            string `position:"Query" name:"CasterId"`
	Effect              string `position:"Query" name:"Effect"`
	ComponentLayer      string `position:"Query" name:"ComponentLayer"`
	CaptionLayerContent string `position:"Query" name:"CaptionLayerContent"`
	ComponentName       string `position:"Query" name:"ComponentName"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	TextLayerContent    string `position:"Query" name:"TextLayerContent"`
}

func (req *ModifyCasterComponentRequest) Invoke(client *sdk.Client) (resp *ModifyCasterComponentResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterComponent", "live", "")
	resp = &ModifyCasterComponentResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCasterComponentResponse struct {
	responses.BaseResponse
	RequestId   string
	ComponentId string
}
