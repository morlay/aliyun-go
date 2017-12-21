package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterVideoResourceRequest struct {
	ResourceId    string `position:"Query" name:"ResourceId"`
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int    `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
	Version       string `position:"Query" name:"Version"`
}

func (r ModifyCasterVideoResourceRequest) Invoke(client *sdk.Client) (response *ModifyCasterVideoResourceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCasterVideoResourceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterVideoResource", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCasterVideoResourceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCasterVideoResourceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCasterVideoResourceResponse struct {
	RequestId  string
	CasterId   string
	ResourceId string
}
