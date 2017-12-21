
package sas-api

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type GetAccountProfileRequest struct {
DeviceIdMd5 string `position:"Query" name:"DeviceIdMd.5"`
Carrier int `position:"Query" name:"Carrier"`
Os string `position:"Query" name:"Os"`
Phone string `position:"Query" name:"Phone"`
RequestUrl string `position:"Query" name:"RequestUrl"`
Ip string `position:"Query" name:"Ip"`
UserAgent string `position:"Query" name:"UserAgent"`
ConnectionType int `position:"Query" name:"ConnectionType"`
SensType int `position:"Query" name:"SensType"`
DeviceType int `position:"Query" name:"DeviceType"`
AccessTimestamp int64 `position:"Query" name:"AccessTimestamp"`
BusinessType int `position:"Query" name:"BusinessType"`
}

func (r GetAccountProfileRequest) Invoke(client *sdk.Client) (response *GetAccountProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAccountProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetAccountProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAccountProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.GetAccountProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAccountProfileResponse struct {
Code int
Message string
Success bool
RequestId string
Data GetAccountProfileData
}

type GetAccountProfileData struct {
Ip string
Phone string
IpInfo string
PhoneInfo string
}

