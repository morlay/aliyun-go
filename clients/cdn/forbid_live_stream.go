package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ForbidLiveStreamRequest struct {
	ResumeTime     string `position:"Query" name:"ResumeTime"`
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
}

func (r ForbidLiveStreamRequest) Invoke(client *sdk.Client) (response *ForbidLiveStreamResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ForbidLiveStreamRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ForbidLiveStream", "", "")

	resp := struct {
		*responses.BaseResponse
		ForbidLiveStreamResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ForbidLiveStreamResponse

	err = client.DoAction(&req, &resp)
	return
}

type ForbidLiveStreamResponse struct {
	RequestId string
}
