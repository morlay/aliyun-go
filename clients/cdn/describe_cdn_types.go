package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnTypesRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnTypesRequest) Invoke(client *sdk.Client) (resp *DescribeCdnTypesResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnTypes", "", "")
	resp = &DescribeCdnTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnTypesResponse struct {
	responses.BaseResponse
	RequestId common.String
	CdnTypes  DescribeCdnTypesCdnTypeList
}

type DescribeCdnTypesCdnType struct {
	Type common.String
	Desc common.String
}

type DescribeCdnTypesCdnTypeList []DescribeCdnTypesCdnType

func (list *DescribeCdnTypesCdnTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnTypesCdnType)
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
