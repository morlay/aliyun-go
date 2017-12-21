package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainISPDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainISPDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainISPDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainISPDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainISPData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainISPDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainISPDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainISPDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	Value        DescribeDomainISPDataISPProportionDataList
}

type DescribeDomainISPDataISPProportionData struct {
	ISP             string
	Proportion      string
	IspEname        string
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

type DescribeDomainISPDataISPProportionDataList []DescribeDomainISPDataISPProportionData

func (list *DescribeDomainISPDataISPProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainISPDataISPProportionData)
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
