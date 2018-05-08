package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDevicesByAccountRequest struct {
	requests.RpcRequest
	AppKey  int64  `position:"Query" name:"AppKey"`
	Account string `position:"Query" name:"Account"`
}

func (req *QueryDevicesByAccountRequest) Invoke(client *sdk.Client) (resp *QueryDevicesByAccountResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryDevicesByAccount", "", "")
	resp = &QueryDevicesByAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDevicesByAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
	DeviceIds QueryDevicesByAccountDeviceIdList
}

type QueryDevicesByAccountDeviceIdList []common.String

func (list *QueryDevicesByAccountDeviceIdList) UnmarshalJSON(data []byte) error {
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
