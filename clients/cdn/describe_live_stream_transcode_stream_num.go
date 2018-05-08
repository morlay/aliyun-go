package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	requests.RpcRequest
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamTranscodeStreamNumRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamTranscodeStreamNumResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamTranscodeStreamNum", "", "")
	resp = &DescribeLiveStreamTranscodeStreamNumResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	responses.BaseResponse
	RequestId         common.String
	Total             common.Long
	TranscodedNumber  common.Long
	UntranscodeNumber common.Long
}
