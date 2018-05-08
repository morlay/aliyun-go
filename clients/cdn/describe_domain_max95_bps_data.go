package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainMax95BpsDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainMax95BpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainMax95BpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainMax95BpsData", "", "")
	resp = &DescribeDomainMax95BpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainMax95BpsDataResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DomainName       common.String
	StartTime        common.String
	EndTime          common.String
	Max95Bps         common.String
	DomesticMax95Bps common.String
	OverseasMax95Bps common.String
}
