package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRelayPushBitRateRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamRelayPushBitRateRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRelayPushBitRateResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushBitRate", "", "")
	resp = &DescribeLiveStreamRelayPushBitRateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRelayPushBitRateResponse struct {
	responses.BaseResponse
	RequestId                 string
	RelayPushBitRateModelList DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel struct {
	Time          string
	VedioFrame    string
	VedioTimstamp string
	AudioFrame    string
	AudioTimstamp string
	RelayDomain   string
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList []DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel

func (list *DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel)
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
