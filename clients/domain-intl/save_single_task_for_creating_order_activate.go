package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	requests.RpcRequest
	SubscriptionDuration    int    `position:"Query" name:"SubscriptionDuration"`
	PermitPremiumActivation string `position:"Query" name:"PermitPremiumActivation"`
	UserClientIp            string `position:"Query" name:"UserClientIp"`
	DomainName              string `position:"Query" name:"DomainName"`
	RegistrantProfileId     int64  `position:"Query" name:"RegistrantProfileId"`
	EnableDomainProxy       string `position:"Query" name:"EnableDomainProxy"`
	Lang                    string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingOrderActivateRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderActivateResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingOrderActivate", "domain", "")
	resp = &SaveSingleTaskForCreatingOrderActivateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
