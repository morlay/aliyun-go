package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamTranscodeInfoRequest struct {
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	DomainTranscodeName string `position:"Query" name:"DomainTranscodeName"`
}

func (r DescribeLiveStreamTranscodeInfoRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamTranscodeInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamTranscodeInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamTranscodeInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamTranscodeInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamTranscodeInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	RequestId           string
	DomainTranscodeList DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp      string
	TranscodeName     string
	TranscodeTemplate string
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
