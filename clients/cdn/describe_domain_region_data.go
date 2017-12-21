package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRegionDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainRegionDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainRegionDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainRegionDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRegionData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainRegionDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainRegionDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainRegionDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	Value        DescribeDomainRegionDataRegionProportionDataList
}

type DescribeDomainRegionDataRegionProportionData struct {
	Region          string
	Proportion      string
	RegionEname     string
	AvgObjectSize   string
	AvgResponseTime string
	Bps             string
	ByteHitRate     string
	Qps             string
	ReqErrRate      string
	ReqHitRate      string
	AvgResponseRate string
	TotalBytes      string
	BytesProportion string
	TotalQuery      string
}

type DescribeDomainRegionDataRegionProportionDataList []DescribeDomainRegionDataRegionProportionData

func (list *DescribeDomainRegionDataRegionProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRegionDataRegionProportionData)
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
