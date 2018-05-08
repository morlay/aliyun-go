package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCACertificatesRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
}

func (req *DescribeCACertificatesRequest) Invoke(client *sdk.Client) (resp *DescribeCACertificatesResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeCACertificates", "slb", "")
	resp = &DescribeCACertificatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCACertificatesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	CACertificates DescribeCACertificatesCACertificateList
}

type DescribeCACertificatesCACertificate struct {
	RegionId          common.String
	CACertificateId   common.String
	CACertificateName common.String
	Fingerprint       common.String
	ResourceGroupId   common.String
	CreateTime        common.String
	CreateTimeStamp   common.Long
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
