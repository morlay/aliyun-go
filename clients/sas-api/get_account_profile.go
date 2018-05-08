package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAccountProfileRequest struct {
	requests.RpcRequest
	DeviceIdMd5     string `position:"Query" name:"DeviceIdMd.5"`
	Carrier         int    `position:"Query" name:"Carrier"`
	Os              string `position:"Query" name:"Os"`
	Phone           string `position:"Query" name:"Phone"`
	RequestUrl      string `position:"Query" name:"RequestUrl"`
	Ip              string `position:"Query" name:"Ip"`
	UserAgent       string `position:"Query" name:"UserAgent"`
	ConnectionType  int    `position:"Query" name:"ConnectionType"`
	SensType        int    `position:"Query" name:"SensType"`
	DeviceType      int    `position:"Query" name:"DeviceType"`
	AccessTimestamp int64  `position:"Query" name:"AccessTimestamp"`
	BusinessType    int    `position:"Query" name:"BusinessType"`
}

func (req *GetAccountProfileRequest) Invoke(client *sdk.Client) (resp *GetAccountProfileResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetAccountProfile", "", "")
	resp = &GetAccountProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountProfileResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	Success   bool
	RequestId common.String
	Data      GetAccountProfileData
}

type GetAccountProfileData struct {
	Ip        common.String
	Phone     common.String
	IpInfo    common.String
	PhoneInfo common.String
}
