package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCACertificateRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
}

func (r DeleteCACertificateRequest) Invoke(client *sdk.Client) (response *DeleteCACertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCACertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteCACertificate", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCACertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCACertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCACertificateResponse struct {
	RequestId string
}
