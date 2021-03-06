package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDevicesByAliasRequest struct {
	requests.RpcRequest
	Alias  string `position:"Query" name:"Alias"`
	AppKey int64  `position:"Query" name:"AppKey"`
}

func (req *QueryDevicesByAliasRequest) Invoke(client *sdk.Client) (resp *QueryDevicesByAliasResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryDevicesByAlias", "", "")
	resp = &QueryDevicesByAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDevicesByAliasResponse struct {
	responses.BaseResponse
	RequestId common.String
	DeviceIds QueryDevicesByAliasDeviceIdList
}

type QueryDevicesByAliasDeviceIdList []common.String

func (list *QueryDevicesByAliasDeviceIdList) UnmarshalJSON(data []byte) error {
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
