package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	requests.RpcRequest
	OrderRedeemParams *SaveBatchTaskForCreatingOrderRedeemOrderRedeemParamList `position:"Query" type:"Repeated" name:"OrderRedeemParam"`
	UserClientIp      string                                                   `position:"Query" name:"UserClientIp"`
	Lang              string                                                   `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderRedeemRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderRedeemResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForCreatingOrderRedeem", "", "")
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
	RequestId string
	TaskNo    string
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
