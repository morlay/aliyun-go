package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForCreatingOrderRenewRequest struct {
	requests.RpcRequest
	UserClientIp     string                                                 `position:"Query" name:"UserClientIp"`
	OrderRenewParams *SaveBatchTaskForCreatingOrderRenewOrderRenewParamList `position:"Query" type:"Repeated" name:"OrderRenewParam"`
	Lang             string                                                 `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderRenewRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderRenewResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForCreatingOrderRenew", "", "")
	resp = &SaveBatchTaskForCreatingOrderRenewResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForCreatingOrderRenewOrderRenewParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate int64  `name:"CurrentExpirationDate"`
	SubscriptionDuration  int    `name:"SubscriptionDuration"`
}

type SaveBatchTaskForCreatingOrderRenewResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveBatchTaskForCreatingOrderRenewOrderRenewParamList []SaveBatchTaskForCreatingOrderRenewOrderRenewParam

func (list *SaveBatchTaskForCreatingOrderRenewOrderRenewParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderRenewOrderRenewParam)
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
