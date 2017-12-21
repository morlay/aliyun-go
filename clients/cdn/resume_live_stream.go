package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResumeLiveStreamRequest struct {
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
}

func (r ResumeLiveStreamRequest) Invoke(client *sdk.Client) (response *ResumeLiveStreamResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResumeLiveStreamRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ResumeLiveStream", "", "")

	resp := struct {
		*responses.BaseResponse
		ResumeLiveStreamResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResumeLiveStreamResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResumeLiveStreamResponse struct {
	RequestId string
}
