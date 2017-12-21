package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadCACertificateRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CACertificate        string `position:"Query" name:"CACertificate"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UploadCACertificateRequest) Invoke(client *sdk.Client) (resp *UploadCACertificateResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "UploadCACertificate", "slb", "")
	resp = &UploadCACertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadCACertificateResponse struct {
	responses.BaseResponse
	RequestId         string
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
	ResourceGroupId   string
	CreateTime        string
	CreateTimeStamp   int64
}
