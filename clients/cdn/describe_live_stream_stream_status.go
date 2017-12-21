package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamStreamStatusRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamStreamStatusRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamStreamStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamStreamStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamStreamStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamStreamStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamStreamStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamStreamStatusResponse struct {
	RequestId    string
	StreamStatus string
}
