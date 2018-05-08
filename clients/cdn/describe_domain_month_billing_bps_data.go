package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId              common.String
	DomainName             common.String
	StartTime              common.String
	EndTime                common.String
	Bandwidth95Bps         common.Float
	DomesticBandwidth95Bps common.Float
	OverseasBandwidth95Bps common.Float
	MonthavgBps            common.Float
	DomesticMonthavgBps    common.Float
	OverseasMonthavgBps    common.Float
	Month4thBps            common.Float
	DomesticMonth4thBps    common.Float
	OverseasMonth4thBps    common.Float
}
