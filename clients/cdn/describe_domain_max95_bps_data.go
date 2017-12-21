package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainMax95BpsDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainMax95BpsDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainMax95BpsDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainMax95BpsDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainMax95BpsData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainMax95BpsDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainMax95BpsDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainMax95BpsDataResponse struct {
	RequestId        string
	DomainName       string
	StartTime        string
	EndTime          string
	Max95Bps         string
	DomesticMax95Bps string
	OverseasMax95Bps string
}
