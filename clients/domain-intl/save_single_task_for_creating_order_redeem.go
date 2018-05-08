package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForCreatingOrderRedeemRequest struct {
	requests.RpcRequest
	CurrentExpirationDate int64  `position:"Query" name:"CurrentExpirationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	Lang                  string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingOrderRedeemRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderRedeemResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingOrderRedeem", "domain", "")
	resp = &SaveSingleTaskForCreatingOrderRedeemResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderRedeemResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
