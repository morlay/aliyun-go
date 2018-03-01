package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamTranscodeInfoRequest struct {
	requests.RpcRequest
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	DomainTranscodeName string `position:"Query" name:"DomainTranscodeName"`
}

func (req *DescribeLiveStreamTranscodeInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamTranscodeInfoResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamTranscodeInfo", "live", "")
	resp = &DescribeLiveStreamTranscodeInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	responses.BaseResponse
	RequestId           string
	DomainTranscodeList DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp              string
	TranscodeName             string
	TranscodeTemplate         string
	CustomTranscodeParameters DescribeLiveStreamTranscodeInfoCustomTranscodeParameters
}

type DescribeLiveStreamTranscodeInfoCustomTranscodeParameters struct {
	VideoBitrate int
	FPS          int
	Height       int
	Width        int
	TemplateType string
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList []DescribeLiveStreamTranscodeInfoDomainTranscodeInfo

func (list *DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamTranscodeInfoDomainTranscodeInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
