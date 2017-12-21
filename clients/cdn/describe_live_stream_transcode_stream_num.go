package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamTranscodeStreamNumRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamTranscodeStreamNumResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamTranscodeStreamNumRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamTranscodeStreamNum", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamTranscodeStreamNumResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamTranscodeStreamNumResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	RequestId         string
	Total             int64
	TranscodedNumber  int64
	UntranscodeNumber int64
}
