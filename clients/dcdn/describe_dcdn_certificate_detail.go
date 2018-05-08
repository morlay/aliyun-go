package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnCertificateDetailRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CertName      string `position:"Query" name:"CertName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnCertificateDetailRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnCertificateDetailResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnCertificateDetail", "dcdn", "")
	resp = &DescribeDcdnCertificateDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnCertificateDetailResponse struct {
	responses.BaseResponse
	RequestId string
	Cert      string
	Key       string
	CertId    int64
	CertName  string
}
