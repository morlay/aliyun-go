package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	requests.RpcRequest
	OrderRedeemParams *SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList `position:"Query" type:"Repeated" name:"OrderRedeemParam"`
	UserClientIp      string                                                   `position:"Query" name:"UserClientIp"`
	Lang              string                                                   `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderRedeemRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderRedeemResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForCreatingOrderRedeem", "domain", "")
	resp = &SaveBatchTaskForCreatingOrderRedeemResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam struct {
	DomainName            string `name:"DomainName"`
	CurrentExpirationDate int64  `name:"CurrentExpirationDate"`
}

type SaveBatchTaskForCreatingOrderRedeemResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList []SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam

func (list *SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderRedeemOrderRedeemParam)
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
