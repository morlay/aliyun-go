package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRegionDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRegionDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRegionDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRegionData", "", "")
	resp = &DescribeDomainRegionDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRegionDataResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DomainName   common.String
	DataInterval common.String
	StartTime    common.String
	EndTime      common.String
	Value        DescribeDomainRegionDataRegionProportionDataList
}

type DescribeDomainRegionDataRegionProportionData struct {
	Region          common.String
	Proportion      common.String
	RegionEname     common.String
	AvgObjectSize   common.String
	AvgResponseTime common.String
	Bps             common.String
	ByteHitRate     common.String
	Qps             common.String
	ReqErrRate      common.String
	ReqHitRate      common.String
	AvgResponseRate common.String
	TotalBytes      common.String
	BytesProportion common.String
	TotalQuery      common.String
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
