package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenVodServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *OpenVodServiceRequest) Invoke(client *sdk.Client) (resp *OpenVodServiceResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "OpenVodService", "", "")
	resp = &OpenVodServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type OpenVodServiceResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
}
