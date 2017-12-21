
package sas-api

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type GetPhoneProfileRequest struct {
Phone string `position:"Query" name:"Phone"`
SensType int `position:"Query" name:"SensType"`
DataVersion string `position:"Query" name:"DataVersion"`
BusinessType int `position:"Query" name:"BusinessType"`
}

func (r GetPhoneProfileRequest) Invoke(client *sdk.Client) (response *GetPhoneProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPhoneProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetPhoneProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		GetPhoneProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.GetPhoneProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPhoneProfileResponse struct {
Code int
Message string
Success bool
RequestId string
Data GetPhoneProfileData
}

type GetPhoneProfileData struct {
Phone string
Info string
}

