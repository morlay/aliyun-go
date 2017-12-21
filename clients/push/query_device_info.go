package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceInfoRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (r QueryDeviceInfoRequest) Invoke(client *sdk.Client) (response *QueryDeviceInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDeviceInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDeviceInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDeviceInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDeviceInfoResponse struct {
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
}
