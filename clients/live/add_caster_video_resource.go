package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCasterVideoResourceRequest struct {
	requests.RpcRequest
	BeginOffset   int    `position:"Query" name:"BeginOffset"`
	VodUrl        string `position:"Query" name:"VodUrl"`
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	LocationId    string `position:"Query" name:"LocationId"`
	CasterId      string `position:"Query" name:"CasterId"`
	EndOffset     int    `position:"Query" name:"EndOffset"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int    `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
}

func (req *AddCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *AddCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterVideoResource", "live", "")
	resp = &AddCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId  string
	ResourceId string
}
