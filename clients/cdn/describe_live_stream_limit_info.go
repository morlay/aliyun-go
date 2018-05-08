package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamLimitInfoRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	LimitDomain   string `position:"Query" name:"LimitDomain"`
}

func (req *DescribeLiveStreamLimitInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamLimitInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamLimitInfo", "", "")
	resp = &DescribeLiveStreamLimitInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamLimitInfoResponse struct {
	responses.BaseResponse
	RequestId      common.String
	UserLimitLists DescribeLiveStreamLimitInfoUserLimitModeList
}

type DescribeLiveStreamLimitInfoUserLimitMode struct {
	LimitDomain       common.String
	LimitNum          common.String
	LimitTranscodeNum common.String
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
