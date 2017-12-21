package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDnsProductAttributesRequest struct {
	VersionCode string `position:"Query" name:"VersionCode"`
}

func (r DescribeDnsProductAttributesRequest) Invoke(client *sdk.Client) (response *DescribeDnsProductAttributesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDnsProductAttributesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductAttributes", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDnsProductAttributesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDnsProductAttributesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDnsProductAttributesResponse struct {
	RequestId   string
	TtlMinValue string
	TtlMaxValue string
	RecordLines DescribeDnsProductAttributesRecordLineList
	RecordTypes DescribeDnsProductAttributesRecordTypeList
}

type DescribeDnsProductAttributesRecordLine struct {
	LineCode   string
	FatherCode string
	LineName   string
}

type DescribeDnsProductAttributesRecordLineList []DescribeDnsProductAttributesRecordLine

func (list *DescribeDnsProductAttributesRecordLineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductAttributesRecordLine)
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

type DescribeDnsProductAttributesRecordTypeList []string

func (list *DescribeDnsProductAttributesRecordTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
