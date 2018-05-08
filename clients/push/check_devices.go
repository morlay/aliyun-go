package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckDevicesRequest struct {
	requests.RpcRequest
	DeviceIds string `position:"Query" name:"DeviceIds"`
	AppKey    int64  `position:"Query" name:"AppKey"`
}

func (req *CheckDevicesRequest) Invoke(client *sdk.Client) (resp *CheckDevicesResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "CheckDevices", "", "")
	resp = &CheckDevicesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDevicesResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DeviceCheckInfos CheckDevicesDeviceCheckInfoList
}

type CheckDevicesDeviceCheckInfo struct {
	DeviceId  common.String
	Available bool
}

type CheckDevicesDeviceCheckInfoList []CheckDevicesDeviceCheckInfo

func (list *CheckDevicesDeviceCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckDevicesDeviceCheckInfo)
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
