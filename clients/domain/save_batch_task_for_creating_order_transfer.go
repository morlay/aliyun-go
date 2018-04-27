package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForCreatingOrderTransferRequest struct {
	requests.RpcRequest
	OrderTransferParams *SaveBatchTaskForCreatingOrderTransferOrderTransferParamList `position:"Query" type:"Repeated" name:"OrderTransferParam"`
	UserClientIp        string                                                       `position:"Query" name:"UserClientIp"`
	Lang                string                                                       `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderTransferRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderTransferResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForCreatingOrderTransfer", "", "")
	resp = &SaveBatchTaskForCreatingOrderTransferResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForCreatingOrderTransferOrderTransferParam struct {
	DomainName            string `name:"DomainName"`
	AuthorizationCode     string `name:"AuthorizationCode"`
	RegistrantProfileId   int64  `name:"RegistrantProfileId"`
	PermitPremiumTransfer string `name:"PermitPremiumTransfer"`
}

type SaveBatchTaskForCreatingOrderTransferResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForCreatingOrderTransferOrderTransferParamList []SaveBatchTaskForCreatingOrderTransferOrderTransferParam

func (list *SaveBatchTaskForCreatingOrderTransferOrderTransferParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderTransferOrderTransferParam)
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
