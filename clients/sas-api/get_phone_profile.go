package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPhoneProfileRequest struct {
	requests.RpcRequest
	Phone        string `position:"Query" name:"Phone"`
	SensType     int    `position:"Query" name:"SensType"`
	DataVersion  string `position:"Query" name:"DataVersion"`
	BusinessType int    `position:"Query" name:"BusinessType"`
}

func (req *GetPhoneProfileRequest) Invoke(client *sdk.Client) (resp *GetPhoneProfileResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetPhoneProfile", "", "")
	resp = &GetPhoneProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPhoneProfileResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	Success   bool
	RequestId string
	Data      GetPhoneProfileData
}

type GetPhoneProfileData struct {
	Phone string
	Info  string
}
