package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TransferInCheckMailTokenRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
}

func (req *TransferInCheckMailTokenRequest) Invoke(client *sdk.Client) (resp *TransferInCheckMailTokenResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "TransferInCheckMailToken", "", "")
	resp = &TransferInCheckMailTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransferInCheckMailTokenResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SuccessList TransferInCheckMailTokenSuccessListList
	FailList    TransferInCheckMailTokenFailListList
}

type TransferInCheckMailTokenSuccessListList []common.String

func (list *TransferInCheckMailTokenSuccessListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type TransferInCheckMailTokenFailListList []common.String

func (list *TransferInCheckMailTokenFailListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
