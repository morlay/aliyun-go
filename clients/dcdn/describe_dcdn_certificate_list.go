package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	CertificateListModel DescribeDcdnCertificateListCertificateListModel
}

type DescribeDcdnCertificateListCertificateListModel struct {
	Count    common.Integer
	CertList DescribeDcdnCertificateListCertList
}

type DescribeDcdnCertificateListCert struct {
	CertName    common.String
	CertId      common.Long
	Fingerprint common.String
	Common      common.String
	Issuer      common.String
	LastTime    common.Long
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
