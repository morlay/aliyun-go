package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCACertificatesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
}

func (r DescribeCACertificatesRequest) Invoke(client *sdk.Client) (response *DescribeCACertificatesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCACertificatesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeCACertificates", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCACertificatesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCACertificatesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCACertificatesResponse struct {
	RequestId      string
	CACertificates DescribeCACertificatesCACertificateList
}

type DescribeCACertificatesCACertificate struct {
	RegionId          string
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
	ResourceGroupId   string
	CreateTime        string
	CreateTimeStamp   int64
}

type DescribeCACertificatesCACertificateList []DescribeCACertificatesCACertificate

func (list *DescribeCACertificatesCACertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCACertificatesCACertificate)
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
