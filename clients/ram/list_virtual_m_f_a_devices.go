package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListVirtualMFADevicesRequest struct {
}

func (r ListVirtualMFADevicesRequest) Invoke(client *sdk.Client) (response *ListVirtualMFADevicesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListVirtualMFADevicesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListVirtualMFADevices", "", "")

	resp := struct {
		*responses.BaseResponse
		ListVirtualMFADevicesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListVirtualMFADevicesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListVirtualMFADevicesResponse struct {
	RequestId         string
	VirtualMFADevices ListVirtualMFADevicesVirtualMFADeviceList
}

type ListVirtualMFADevicesVirtualMFADevice struct {
	SerialNumber string
	ActivateDate string
	User         ListVirtualMFADevicesUser
}

type ListVirtualMFADevicesUser struct {
	UserId      string
	UserName    string
	DisplayName string
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
