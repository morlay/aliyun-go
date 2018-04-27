package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterVideoResourceRequest struct {
	requests.RpcRequest
	ResourceId    string `position:"Query" name:"ResourceId"`
	BeginOffset   int    `position:"Query" name:"BeginOffset"`
	VodUrl        string `position:"Query" name:"VodUrl"`
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	CasterId      string `position:"Query" name:"CasterId"`
	EndOffset     int    `position:"Query" name:"EndOffset"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int    `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
}

func (req *ModifyCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *ModifyCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterVideoResource", "live", "")
	resp = &ModifyCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId  string
	CasterId   string
	ResourceId string
}
