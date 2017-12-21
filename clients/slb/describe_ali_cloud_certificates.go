package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAliCloudCertificatesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeAliCloudCertificatesRequest) Invoke(client *sdk.Client) (response *DescribeAliCloudCertificatesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAliCloudCertificatesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeAliCloudCertificates", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAliCloudCertificatesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAliCloudCertificatesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAliCloudCertificatesResponse struct {
	RequestId            string
	AliCloudCertificates DescribeAliCloudCertificatesAliCloudCertificateList
}

type DescribeAliCloudCertificatesAliCloudCertificate struct {
	AliCloudCertificateId   string
	AliCloudCertificateName string
	Fingerprint             string
	DomainName              string
	Issuer                  string
}

type DescribeAliCloudCertificatesAliCloudCertificateList []DescribeAliCloudCertificatesAliCloudCertificate

func (list *DescribeAliCloudCertificatesAliCloudCertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAliCloudCertificatesAliCloudCertificate)
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
