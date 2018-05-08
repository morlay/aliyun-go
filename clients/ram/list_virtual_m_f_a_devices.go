package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListVirtualMFADevicesRequest struct {
	requests.RpcRequest
}

func (req *ListVirtualMFADevicesRequest) Invoke(client *sdk.Client) (resp *ListVirtualMFADevicesResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListVirtualMFADevices", "", "")
	resp = &ListVirtualMFADevicesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListVirtualMFADevicesResponse struct {
	responses.BaseResponse
	RequestId         common.String
	VirtualMFADevices ListVirtualMFADevicesVirtualMFADeviceList
}

type ListVirtualMFADevicesVirtualMFADevice struct {
	SerialNumber common.String
	ActivateDate common.String
	User         ListVirtualMFADevicesUser
}

type ListVirtualMFADevicesUser struct {
	UserId      common.String
	UserName    common.String
	DisplayName common.String
}

type ListVirtualMFADevicesVirtualMFADeviceList []ListVirtualMFADevicesVirtualMFADevice

func (list *ListVirtualMFADevicesVirtualMFADeviceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVirtualMFADevicesVirtualMFADevice)
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
