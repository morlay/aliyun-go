package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCACertificateRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
}

func (req *DeleteCACertificateRequest) Invoke(client *sdk.Client) (resp *DeleteCACertificateResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteCACertificate", "slb", "")
	resp = &DeleteCACertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCACertificateResponse struct {
	responses.BaseResponse
	RequestId string
}
