package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCreatingOrderRenewRequest struct {
	requests.RpcRequest
	SubscriptionDuration  int    `position:"Query" name:"SubscriptionDuration"`
	CurrentExpirationDate int64  `position:"Query" name:"CurrentExpirationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	Lang                  string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingOrderRenewRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderRenewResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForCreatingOrderRenew", "", "")
	resp = &SaveSingleTaskForCreatingOrderRenewResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderRenewResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
