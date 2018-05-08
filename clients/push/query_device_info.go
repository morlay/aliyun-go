package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	DeviceInfo QueryDeviceInfoDeviceInfo
}

type QueryDeviceInfoDeviceInfo struct {
	DeviceId       common.String
	DeviceType     common.String
	Account        common.String
	DeviceToken    common.String
	Tags           common.String
	Alias          common.String
	LastOnlineTime common.String
	Online         bool
	PhoneNumber    common.String
	PushEnabled    bool
}
