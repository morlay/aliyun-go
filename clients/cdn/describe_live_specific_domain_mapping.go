package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveSpecificDomainMappingRequest struct {
	requests.RpcRequest
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveSpecificDomainMappingRequest) Invoke(client *sdk.Client) (resp *DescribeLiveSpecificDomainMappingResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveSpecificDomainMapping", "", "")
	resp = &DescribeLiveSpecificDomainMappingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveSpecificDomainMappingResponse struct {
	responses.BaseResponse
	RequestId           string
	DomainMappingModels DescribeLiveSpecificDomainMappingDomainMappingModelList
}

type DescribeLiveSpecificDomainMappingDomainMappingModel struct {
	PushDomain string
	PullDomain string
}

type DescribeLiveSpecificDomainMappingDomainMappingModelList []DescribeLiveSpecificDomainMappingDomainMappingModel

func (list *DescribeLiveSpecificDomainMappingDomainMappingModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSpecificDomainMappingDomainMappingModel)
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
