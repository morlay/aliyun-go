package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetServerCertificateNameRequest struct {
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerCertificateId   string `position:"Query" name:"ServerCertificateId"`
	ServerCertificateName string `position:"Query" name:"ServerCertificateName"`
	Tags                  string `position:"Query" name:"Tags"`
}

func (r SetServerCertificateNameRequest) Invoke(client *sdk.Client) (response *SetServerCertificateNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetServerCertificateNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetServerCertificateName", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetServerCertificateNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetServerCertificateNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetServerCertificateNameResponse struct {
	RequestId string
}
