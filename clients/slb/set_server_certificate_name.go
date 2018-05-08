package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetServerCertificateNameRequest struct {
	requests.RpcRequest
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId   string `position:"Query" name:"ServerCertificateId"`
	ServerCertificateName string `position:"Query" name:"ServerCertificateName"`
	Tags                  string `position:"Query" name:"Tags"`
}

func (req *SetServerCertificateNameRequest) Invoke(client *sdk.Client) (resp *SetServerCertificateNameResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetServerCertificateName", "slb", "")
	resp = &SetServerCertificateNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetServerCertificateNameResponse struct {
	responses.BaseResponse
	RequestId common.String
}
