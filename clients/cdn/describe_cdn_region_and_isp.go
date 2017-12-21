package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnRegionAndIspRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCdnRegionAndIspRequest) Invoke(client *sdk.Client) (response *DescribeCdnRegionAndIspResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnRegionAndIspRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnRegionAndIsp", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnRegionAndIspResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnRegionAndIspResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnRegionAndIspResponse struct {
	RequestId string
	Regions   DescribeCdnRegionAndIspRegionList
	Isps      DescribeCdnRegionAndIspIspList
}

type DescribeCdnRegionAndIspRegion struct {
	NameZh string
	NameEn string
}

type DescribeCdnRegionAndIspIsp struct {
	NameZh string
	NameEn string
}

type DescribeCdnRegionAndIspRegionList []DescribeCdnRegionAndIspRegion

func (list *DescribeCdnRegionAndIspRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspRegion)
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

type DescribeCdnRegionAndIspIspList []DescribeCdnRegionAndIspIsp

func (list *DescribeCdnRegionAndIspIspList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspIsp)
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
