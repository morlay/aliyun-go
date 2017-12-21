package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadServerCertificateRequest struct {
	requests.RpcRequest
	Access_key_id           string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	AliCloudCertificateName string `position:"Query" name:"AliCloudCertificateName"`
	AliCloudCertificateId   string `position:"Query" name:"AliCloudCertificateId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Tags                    string `position:"Query" name:"Tags"`
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ResourceGroupId         string `position:"Query" name:"ResourceGroupId"`
	ServerCertificateName   string `position:"Query" name:"ServerCertificateName"`
}

func (req *UploadServerCertificateRequest) Invoke(client *sdk.Client) (resp *UploadServerCertificateResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "UploadServerCertificate", "slb", "")
	resp = &UploadServerCertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadServerCertificateResponse struct {
	responses.BaseResponse
	RequestId               string
	ServerCertificateId     string
	Fingerprint             string
	ServerCertificateName   string
	RegionId                string
	RegionIdAlias           string
	AliCloudCertificateId   string
	AliCloudCertificateName string
	IsAliCloudCertificate   int
	ResourceGroupId         string
	CreateTime              string
	CreateTimeStamp         int64
}
