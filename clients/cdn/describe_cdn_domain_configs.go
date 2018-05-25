package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnDomainConfigsRequest struct {
	requests.RpcRequest
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnDomainConfigsRequest) Invoke(client *sdk.Client) (resp *DescribeCdnDomainConfigsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainConfigs", "", "")
	resp = &DescribeCdnDomainConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnDomainConfigsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	DomainConfigs DescribeCdnDomainConfigsDomainConfigList
}

type DescribeCdnDomainConfigsDomainConfig struct {
	FunctionName common.String
	ConfigId     common.String
	Status       common.String
	FunctionArgs DescribeCdnDomainConfigsFunctionArgList
}

type DescribeCdnDomainConfigsFunctionArg struct {
	ArgName  common.String
	ArgValue common.String
}

type DescribeCdnDomainConfigsDomainConfigList []DescribeCdnDomainConfigsDomainConfig

func (list *DescribeCdnDomainConfigsDomainConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainConfigsDomainConfig)
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

type DescribeCdnDomainConfigsFunctionArgList []DescribeCdnDomainConfigsFunctionArg

func (list *DescribeCdnDomainConfigsFunctionArgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainConfigsFunctionArg)
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
