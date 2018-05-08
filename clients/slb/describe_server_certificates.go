package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeServerCertificatesRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId  string `position:"Query" name:"ServerCertificateId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeServerCertificatesRequest) Invoke(client *sdk.Client) (resp *DescribeServerCertificatesResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeServerCertificates", "slb", "")
	resp = &DescribeServerCertificatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeServerCertificatesResponse struct {
	responses.BaseResponse
	RequestId          common.String
	ServerCertificates DescribeServerCertificatesServerCertificateList
}

type DescribeServerCertificatesServerCertificate struct {
	ServerCertificateId     common.String
	Fingerprint             common.String
	ServerCertificateName   common.String
	RegionId                common.String
	RegionIdAlias           common.String
	AliCloudCertificateId   common.String
	AliCloudCertificateName common.String
	IsAliCloudCertificate   common.Integer
	ResourceGroupId         common.String
	CreateTime              common.String
	CreateTimeStamp         common.Long
}

type DescribeServerCertificatesServerCertificateList []DescribeServerCertificatesServerCertificate

func (list *DescribeServerCertificatesServerCertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerCertificatesServerCertificate)
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
