package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCreatingOrderRedeemRequest struct {
	requests.RpcRequest
	CurrentExpirationDate int64  `position:"Query" name:"CurrentExpirationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	Lang                  string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingOrderRedeemRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderRedeemResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForCreatingOrderRedeem", "", "")
	resp = &SaveSingleTaskForCreatingOrderRedeemResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderRedeemResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
