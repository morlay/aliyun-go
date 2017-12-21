package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCasterVideoResourceRequest struct {
	requests.RpcRequest
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	LocationId    string `position:"Query" name:"LocationId"`
	CasterId      string `position:"Query" name:"CasterId"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int    `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *AddCasterVideoResourceRequest) Invoke(client *sdk.Client) (resp *AddCasterVideoResourceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterVideoResource", "", "")
	resp = &AddCasterVideoResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterVideoResourceResponse struct {
	responses.BaseResponse
	RequestId  string
	ResourceId string
}
