package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDnsProductAttributesRequest struct {
	requests.RpcRequest
	VersionCode string `position:"Query" name:"VersionCode"`
}

func (req *DescribeDnsProductAttributesRequest) Invoke(client *sdk.Client) (resp *DescribeDnsProductAttributesResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductAttributes", "", "")
	resp = &DescribeDnsProductAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDnsProductAttributesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TtlMinValue common.String
	TtlMaxValue common.String
	RecordLines DescribeDnsProductAttributesRecordLineList
	RecordTypes DescribeDnsProductAttributesRecordTypeList
}

type DescribeDnsProductAttributesRecordLine struct {
	LineCode   common.String
	FatherCode common.String
	LineName   common.String
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

type DescribeDnsProductAttributesRecordTypeList []common.String

func (list *DescribeDnsProductAttributesRecordTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
