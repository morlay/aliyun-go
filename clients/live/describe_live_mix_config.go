package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveMixConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveMixConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveMixConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveMixConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMixConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveMixConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveMixConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveMixConfigResponse struct {
	RequestId     string
	MixConfigList DescribeLiveMixConfigMixConfigList
}

type DescribeLiveMixConfigMixConfig struct {
	DomainName string
	AppName    string
	Template   string
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
