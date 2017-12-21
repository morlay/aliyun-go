package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MigrateDomainToHttpsDeliveryRequest struct {
	requests.RpcRequest
	PrivateKey        string `position:"Query" name:"PrivateKey"`
	ServerCertificate string `position:"Query" name:"ServerCertificate"`
	SecurityToken     string `position:"Query" name:"SecurityToken"`
	OwnerAccount      string `position:"Query" name:"OwnerAccount"`
	DomainName        string `position:"Query" name:"DomainName"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
}

func (req *MigrateDomainToHttpsDeliveryRequest) Invoke(client *sdk.Client) (resp *MigrateDomainToHttpsDeliveryResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "MigrateDomainToHttpsDelivery", "", "")
	resp = &MigrateDomainToHttpsDeliveryResponse{}
	err = client.DoAction(req, resp)
	return
}

type MigrateDomainToHttpsDeliveryResponse struct {
	responses.BaseResponse
	RequestId string
}
