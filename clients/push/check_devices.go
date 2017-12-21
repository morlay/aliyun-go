package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDevicesRequest struct {
	DeviceIds string `position:"Query" name:"DeviceIds"`
	AppKey    int64  `position:"Query" name:"AppKey"`
}

func (r CheckDevicesRequest) Invoke(client *sdk.Client) (response *CheckDevicesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDevicesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "CheckDevices", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckDevicesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDevicesResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDevicesResponse struct {
	RequestId        string
	DeviceCheckInfos CheckDevicesDeviceCheckInfoList
}

type CheckDevicesDeviceCheckInfo struct {
	DeviceId  string
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
