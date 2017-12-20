package sas_api

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *SasApiClient) GetPhoneProfile(req *GetPhoneProfileArgs) (resp *GetPhoneProfileResponse, err error) {
	resp = &GetPhoneProfileResponse{}
	err = c.Request("GetPhoneProfile", req, resp)
	return
}

type GetPhoneProfileData struct {
	Phone string
	Info  string
}
type GetPhoneProfileArgs struct {
	Phone        string
	SensType     int
	DataVersion  string
	BusinessType int
}
type GetPhoneProfileResponse struct {
	Code      int
	Message   string
	Success   core.Bool
	RequestId string
	Data      GetPhoneProfileData
}

func (c *SasApiClient) GetAccountProfile(req *GetAccountProfileArgs) (resp *GetAccountProfileResponse, err error) {
	resp = &GetAccountProfileResponse{}
	err = c.Request("GetAccountProfile", req, resp)
	return
}

type GetAccountProfileData struct {
	Ip        string
	Phone     string
	IpInfo    string
	PhoneInfo string
}
type GetAccountProfileArgs struct {
	DeviceIdMd5     string
	Carrier         int
	Os              string
	Phone           string
	RequestUrl      string
	Ip              string
	UserAgent       string
	ConnectionType  int
	SensType        int
	DeviceType      int
	AccessTimestamp int64
	BusinessType    int
}
type GetAccountProfileResponse struct {
	Code      int
	Message   string
	Success   core.Bool
	RequestId string
	Data      GetAccountProfileData
}

func (c *SasApiClient) GetIpProfile(req *GetIpProfileArgs) (resp *GetIpProfileResponse, err error) {
	resp = &GetIpProfileResponse{}
	err = c.Request("GetIpProfile", req, resp)
	return
}

type GetIpProfileData struct {
	Ip   string
	Info string
}
type GetIpProfileArgs struct {
	DeviceIdMd5    string
	Carrier        int
	Os             string
	RequestUrl     string
	Ip             string
	UserAgent      string
	ConnectionType int
	SensType       int
	DeviceType     int
	BusinessType   int
}
type GetIpProfileResponse struct {
	Code      int
	Message   string
	Success   core.Bool
	RequestId string
	Data      GetIpProfileData
}

func (c *SasApiClient) GetInstanceCount(req *GetInstanceCountArgs) (resp *GetInstanceCountResponse, err error) {
	resp = &GetInstanceCountResponse{}
	err = c.Request("GetInstanceCount", req, resp)
	return
}

type GetInstanceCountArgs struct {
	OwnerId int64
}
type GetInstanceCountResponse struct {
	Code      string
	Message   string
	Success   core.Bool
	RequestId string
	Data      int
}
