package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCACertificateNameRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
}

func (req *SetCACertificateNameRequest) Invoke(client *sdk.Client) (resp *SetCACertificateNameResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetCACertificateName", "slb", "")
	resp = &SetCACertificateNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetCACertificateNameResponse struct {
	responses.BaseResponse
	RequestId string
}
