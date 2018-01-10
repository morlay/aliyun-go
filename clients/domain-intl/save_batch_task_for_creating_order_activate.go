package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	requests.RpcRequest
	OrderActivateParams *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList `position:"Query" type:"Repeated" name:"OrderActivateParam"`
	UserClientIp        string                                                       `position:"Query" name:"UserClientIp"`
	Lang                string                                                       `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderActivateRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderActivateResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForCreatingOrderActivate", "domain", "")
	resp = &SaveBatchTaskForCreatingOrderActivateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParam struct {
	DomainName              string `name:"DomainName"`
	SubscriptionDuration    int    `name:"SubscriptionDuration"`
	RegistrantProfileId     int64  `name:"RegistrantProfileId"`
	EnableDomainProxy       string `name:"EnableDomainProxy"`
	PermitPremiumActivation string `name:"PermitPremiumActivation"`
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParamList []SaveBatchTaskForCreatingOrderActivateOrderActivateParam

func (list *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderActivateOrderActivateParam)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
