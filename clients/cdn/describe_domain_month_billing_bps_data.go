package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainMonthBillingBpsDataRequest struct {
	requests.RpcRequest
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	DomainName         string `position:"Query" name:"DomainName"`
	EndTime            string `position:"Query" name:"EndTime"`
	StartTime          string `position:"Query" name:"StartTime"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainMonthBillingBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainMonthBillingBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainMonthBillingBpsData", "", "")
	resp = &DescribeDomainMonthBillingBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainMonthBillingBpsDataResponse struct {
	responses.BaseResponse
	RequestId              string
	DomainName             string
	StartTime              string
	EndTime                string
	Bandwidth95Bps         float32
	DomesticBandwidth95Bps float32
	OverseasBandwidth95Bps float32
	MonthavgBps            float32
	DomesticMonthavgBps    float32
	OverseasMonthavgBps    float32
	Month4thBps            float32
	DomesticMonth4thBps    float32
	OverseasMonth4thBps    float32
}
