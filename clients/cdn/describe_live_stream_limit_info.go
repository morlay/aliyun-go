package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamLimitInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	LimitDomain   string `position:"Query" name:"LimitDomain"`
}

func (r DescribeLiveStreamLimitInfoRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamLimitInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamLimitInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamLimitInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamLimitInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamLimitInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamLimitInfoResponse struct {
	RequestId      string
	UserLimitLists DescribeLiveStreamLimitInfoUserLimitModeList
}

type DescribeLiveStreamLimitInfoUserLimitMode struct {
	LimitDomain       string
	LimitNum          string
	LimitTranscodeNum string
}

type DescribeLiveStreamLimitInfoUserLimitModeList []DescribeLiveStreamLimitInfoUserLimitMode

func (list *DescribeLiveStreamLimitInfoUserLimitModeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamLimitInfoUserLimitMode)
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
