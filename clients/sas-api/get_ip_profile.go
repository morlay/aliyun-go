
package sas-api

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type GetIpProfileRequest struct {
DeviceIdMd5 string `position:"Query" name:"DeviceIdMd.5"`
Carrier int `position:"Query" name:"Carrier"`
Os string `position:"Query" name:"Os"`
RequestUrl string `position:"Query" name:"RequestUrl"`
Ip string `position:"Query" name:"Ip"`
UserAgent string `position:"Query" name:"UserAgent"`
ConnectionType int `position:"Query" name:"ConnectionType"`
SensType int `position:"Query" name:"SensType"`
DeviceType int `position:"Query" name:"DeviceType"`
BusinessType int `position:"Query" name:"BusinessType"`
}

func (r GetIpProfileRequest) Invoke(client *sdk.Client) (response *GetIpProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetIpProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetIpProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		GetIpProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.GetIpProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetIpProfileResponse struct {
Code int
Message string
Success bool
RequestId string
Data GetIpProfileData
}

type GetIpProfileData struct {
Ip string
Info string
}

