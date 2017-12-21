package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetIpProfileRequest struct {
	requests.RpcRequest
	DeviceIdMd5    string `position:"Query" name:"DeviceIdMd.5"`
	Carrier        int    `position:"Query" name:"Carrier"`
	Os             string `position:"Query" name:"Os"`
	RequestUrl     string `position:"Query" name:"RequestUrl"`
	Ip             string `position:"Query" name:"Ip"`
	UserAgent      string `position:"Query" name:"UserAgent"`
	ConnectionType int    `position:"Query" name:"ConnectionType"`
	SensType       int    `position:"Query" name:"SensType"`
	DeviceType     int    `position:"Query" name:"DeviceType"`
	BusinessType   int    `position:"Query" name:"BusinessType"`
}

func (req *GetIpProfileRequest) Invoke(client *sdk.Client) (resp *GetIpProfileResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetIpProfile", "", "")
	resp = &GetIpProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetIpProfileResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	Success   bool
	RequestId string
	Data      GetIpProfileData
}

type GetIpProfileData struct {
	Ip   string
	Info string
}
