package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainCertificateInfoRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainCertificateInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDomainCertificateInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCertificateInfo", "", "")
	resp = &DescribeDomainCertificateInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainCertificateInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	CertInfos DescribeDomainCertificateInfoCertInfoList
}

type DescribeDomainCertificateInfoCertInfo struct {
	DomainName              common.String
	CertName                common.String
	CertDomainName          common.String
	CertExpireTime          common.String
	CertLife                common.String
	CertOrg                 common.String
	CertType                common.String
	ServerCertificateStatus common.String
	Status                  common.String
}

type DescribeDomainCertificateInfoCertInfoList []DescribeDomainCertificateInfoCertInfo

func (list *DescribeDomainCertificateInfoCertInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCertificateInfoCertInfo)
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
