package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceInfoRequest struct {
	requests.RpcRequest
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (req *QueryDeviceInfoRequest) Invoke(client *sdk.Client) (resp *QueryDeviceInfoResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceInfo", "", "")
	resp = &QueryDeviceInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceInfoResponse struct {
	responses.BaseResponse
	RequestId  string
	DeviceInfo QueryDeviceInfoDeviceInfo
}

type QueryDeviceInfoDeviceInfo struct {
	DeviceId       string
	DeviceType     string
	Account        string
	DeviceToken    string
	Tags           string
	Alias          string
	LastOnlineTime string
	Online         bool
	PhoneNumber    string
	PushEnabled    bool
}
