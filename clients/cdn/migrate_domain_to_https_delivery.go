package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MigrateDomainToHttpsDeliveryRequest struct {
	PrivateKey        string `position:"Query" name:"PrivateKey"`
	ServerCertificate string `position:"Query" name:"ServerCertificate"`
	SecurityToken     string `position:"Query" name:"SecurityToken"`
	OwnerAccount      string `position:"Query" name:"OwnerAccount"`
	DomainName        string `position:"Query" name:"DomainName"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
}

func (r MigrateDomainToHttpsDeliveryRequest) Invoke(client *sdk.Client) (response *MigrateDomainToHttpsDeliveryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		MigrateDomainToHttpsDeliveryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "MigrateDomainToHttpsDelivery", "", "")

	resp := struct {
		*responses.BaseResponse
		MigrateDomainToHttpsDeliveryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.MigrateDomainToHttpsDeliveryResponse

	err = client.DoAction(&req, &resp)
	return
}

type MigrateDomainToHttpsDeliveryResponse struct {
	RequestId string
}
