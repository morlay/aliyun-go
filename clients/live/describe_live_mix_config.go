package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveMixConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveMixConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveMixConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMixConfig", "live", "")
	resp = &DescribeLiveMixConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveMixConfigResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MixConfigList DescribeLiveMixConfigMixConfigList
}

type DescribeLiveMixConfigMixConfig struct {
	DomainName common.String
	AppName    common.String
	Template   common.String
}

type DescribeLiveMixConfigMixConfigList []DescribeLiveMixConfigMixConfig

func (list *DescribeLiveMixConfigMixConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveMixConfigMixConfig)
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
