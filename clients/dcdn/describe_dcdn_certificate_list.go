package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnCertificateListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnCertificateListRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnCertificateListResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnCertificateList", "dcdn", "")
	resp = &DescribeDcdnCertificateListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnCertificateListResponse struct {
	responses.BaseResponse
	RequestId            string
	CertificateListModel DescribeDcdnCertificateListCertificateListModel
}

type DescribeDcdnCertificateListCertificateListModel struct {
	Count    int
	CertList DescribeDcdnCertificateListCertList
}

type DescribeDcdnCertificateListCert struct {
	CertName    string
	CertId      int64
	Fingerprint string
	Common      string
	Issuer      string
	LastTime    int64
}

type DescribeDcdnCertificateListCertList []DescribeDcdnCertificateListCert

func (list *DescribeDcdnCertificateListCertList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnCertificateListCert)
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
