package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadCACertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CACertificate        string `position:"Query" name:"CACertificate"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UploadCACertificateRequest) Invoke(client *sdk.Client) (response *UploadCACertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadCACertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "UploadCACertificate", "slb", "")

	resp := struct {
		*responses.BaseResponse
		UploadCACertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadCACertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadCACertificateResponse struct {
	RequestId         string
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
	ResourceGroupId   string
	CreateTime        string
	CreateTimeStamp   int64
}
