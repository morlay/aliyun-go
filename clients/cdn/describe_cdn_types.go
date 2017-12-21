package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnTypesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCdnTypesRequest) Invoke(client *sdk.Client) (response *DescribeCdnTypesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnTypesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnTypes", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnTypesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnTypesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnTypesResponse struct {
	RequestId string
	CdnTypes  DescribeCdnTypesCdnTypeList
}

type DescribeCdnTypesCdnType struct {
	Type string
	Desc string
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
