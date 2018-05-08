package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainConfigsRequest struct {
	requests.RpcRequest
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnDomainConfigsRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainConfigsResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainConfigs", "dcdn", "")
	resp = &DescribeDcdnDomainConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainConfigsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	DomainConfigs DescribeDcdnDomainConfigsDomainConfigList
}

type DescribeDcdnDomainConfigsDomainConfig struct {
	FunctionName common.String
	ConfigId     common.String
	Status       common.String
	FunctionArgs DescribeDcdnDomainConfigsFunctionArgList
}

type DescribeDcdnDomainConfigsFunctionArg struct {
	ArgName  common.String
	ArgValue common.String
}

type DescribeDcdnDomainConfigsDomainConfigList []DescribeDcdnDomainConfigsDomainConfig

func (list *DescribeDcdnDomainConfigsDomainConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainConfigsDomainConfig)
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

type DescribeDcdnDomainConfigsFunctionArgList []DescribeDcdnDomainConfigsFunctionArg

func (list *DescribeDcdnDomainConfigsFunctionArgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainConfigsFunctionArg)
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
