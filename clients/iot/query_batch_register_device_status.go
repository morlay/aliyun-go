package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryBatchRegisterDeviceStatusRequest struct {
	requests.RpcRequest
	ApplyId    int64  `position:"Query" name:"ApplyId"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryBatchRegisterDeviceStatusRequest) Invoke(client *sdk.Client) (resp *QueryBatchRegisterDeviceStatusResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryBatchRegisterDeviceStatus", "", "")
	resp = &QueryBatchRegisterDeviceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryBatchRegisterDeviceStatusResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryBatchRegisterDeviceStatusData
}

type QueryBatchRegisterDeviceStatusData struct {
	Status      common.String
	ValidList   QueryBatchRegisterDeviceStatusValidListList
	InvalidList QueryBatchRegisterDeviceStatusInvalidListList
}

type QueryBatchRegisterDeviceStatusValidListList []common.String

func (list *QueryBatchRegisterDeviceStatusValidListList) UnmarshalJSON(data []byte) error {
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

type QueryBatchRegisterDeviceStatusInvalidListList []common.String

func (list *QueryBatchRegisterDeviceStatusInvalidListList) UnmarshalJSON(data []byte) error {
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
