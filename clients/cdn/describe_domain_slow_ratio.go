package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainSlowRatioRequest struct {
	StartTime     string `position:"Query" name:"StartTime"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PageSize      int    `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DescribeDomainSlowRatioRequest) Invoke(client *sdk.Client) (response *DescribeDomainSlowRatioResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainSlowRatioRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSlowRatio", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainSlowRatioResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainSlowRatioResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainSlowRatioResponse struct {
	EndTime                  string
	DataInterval             int
	PageNumber               int
	PageSize                 int
	TotalCount               int
	StartTime                string
	SlowRatioDataPerInterval DescribeDomainSlowRatioSlowRatioDataList
}

type DescribeDomainSlowRatioSlowRatioData struct {
	TotalUsers   int
	SlowUsers    int
	SlowRatio    float32
	RegionNameZh string
	RegionNameEn string
	IspNameZh    string
	IspNameEn    string
	Time         string
}

type DescribeDomainSlowRatioSlowRatioDataList []DescribeDomainSlowRatioSlowRatioData

func (list *DescribeDomainSlowRatioSlowRatioDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSlowRatioSlowRatioData)
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
