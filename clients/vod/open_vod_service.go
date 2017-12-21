package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenVodServiceRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r OpenVodServiceRequest) Invoke(client *sdk.Client) (response *OpenVodServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OpenVodServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "OpenVodService", "", "")

	resp := struct {
		*responses.BaseResponse
		OpenVodServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OpenVodServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type OpenVodServiceResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
}
