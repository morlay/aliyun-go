package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteServerCertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId  string `position:"Query" name:"ServerCertificateId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DeleteServerCertificateRequest) Invoke(client *sdk.Client) (response *DeleteServerCertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteServerCertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteServerCertificate", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteServerCertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteServerCertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteServerCertificateResponse struct {
	RequestId string
}
