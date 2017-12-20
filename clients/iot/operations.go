package iot

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *IotClient) RRpc(req *RRpcArgs) (resp *RRpcResponse, err error) {
	resp = &RRpcResponse{}
	err = c.Request("RRpc", req, resp)
	return
}

type RRpcArgs struct {
	RequestBase64Byte string
	DeviceName        string
	ProductKey        string
	Timeout           int
}
type RRpcResponse struct {
	RequestId         string
	Success           core.Bool
	ErrorMessage      string
	RrpcCode          string
	PayloadBase64Byte string
	MessageId         int64
}

func (c *IotClient) Pub(req *PubArgs) (resp *PubResponse, err error) {
	resp = &PubResponse{}
	err = c.Request("Pub", req, resp)
	return
}

type PubArgs struct {
	TopicFullName  string
	Qos            int
	MessageContent string
	ProductKey     string
}
type PubResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	MessageId    string
}

func (c *IotClient) QueryPageByApplyId(req *QueryPageByApplyIdArgs) (resp *QueryPageByApplyIdResponse, err error) {
	resp = &QueryPageByApplyIdResponse{}
	err = c.Request("QueryPageByApplyId", req, resp)
	return
}

type QueryPageByApplyIdApplyDeviceInfo struct {
	DeviceId     string
	DeviceName   string
	DeviceSecret string
}
type QueryPageByApplyIdArgs struct {
	ApplyId     int64
	PageSize    int
	CurrentPage int
}
type QueryPageByApplyIdResponse struct {
	RequestId       string
	Success         core.Bool
	ErrorMessage    string
	PageSize        int
	Page            int
	PageCount       int
	Total           int
	ApplyDeviceList QueryPageByApplyIdApplyDeviceInfoList
}

type QueryPageByApplyIdApplyDeviceInfoList []QueryPageByApplyIdApplyDeviceInfo

func (list *QueryPageByApplyIdApplyDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPageByApplyIdApplyDeviceInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *IotClient) QueryDeviceProp(req *QueryDevicePropArgs) (resp *QueryDevicePropResponse, err error) {
	resp = &QueryDevicePropResponse{}
	err = c.Request("QueryDeviceProp", req, resp)
	return
}

type QueryDevicePropArgs struct {
	DeviceName string
	ProductKey string
}
type QueryDevicePropResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	Props        string
}

func (c *IotClient) CreateProduct(req *CreateProductArgs) (resp *CreateProductResponse, err error) {
	resp = &CreateProductResponse{}
	err = c.Request("CreateProduct", req, resp)
	return
}

type CreateProductProductInfo struct {
	ProductKey    string
	ProductName   string
	CatId         int64
	CreateUserId  int64
	ProductDesc   string
	FromSource    string
	ExtProps      string
	GmtCreate     string
	GmtModified   string
	ProductSecret string
}
type CreateProductArgs struct {
	CatId          int64
	Name           string
	ExtProps       string
	SecurityPolicy string
	Desc           string
}
type CreateProductResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	ProductInfo  CreateProductProductInfo
}

func (c *IotClient) QueryDeviceByName(req *QueryDeviceByNameArgs) (resp *QueryDeviceByNameResponse, err error) {
	resp = &QueryDeviceByNameResponse{}
	err = c.Request("QueryDeviceByName", req, resp)
	return
}

type QueryDeviceByNameDeviceInfo struct {
	DeviceId     string
	DeviceSecret string
	ProductKey   string
	DeviceStatus string
	DeviceName   string
	DeviceType   string
	GmtCreate    string
	GmtModified  string
}
type QueryDeviceByNameArgs struct {
	DeviceName string
	ProductKey string
}
type QueryDeviceByNameResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	DeviceInfo   QueryDeviceByNameDeviceInfo
}

func (c *IotClient) GetDeviceShadow(req *GetDeviceShadowArgs) (resp *GetDeviceShadowResponse, err error) {
	resp = &GetDeviceShadowResponse{}
	err = c.Request("GetDeviceShadow", req, resp)
	return
}

type GetDeviceShadowArgs struct {
	ShadowMessage string
	DeviceName    string
	ProductKey    string
}
type GetDeviceShadowResponse struct {
	RequestId     string
	Success       core.Bool
	ErrorMessage  string
	ShadowMessage string
}

func (c *IotClient) DeleteDeviceProp(req *DeleteDevicePropArgs) (resp *DeleteDevicePropResponse, err error) {
	resp = &DeleteDevicePropResponse{}
	err = c.Request("DeleteDeviceProp", req, resp)
	return
}

type DeleteDevicePropArgs struct {
	DeviceName string
	ProductKey string
	PropKey    string
}
type DeleteDevicePropResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
}

func (c *IotClient) ApplyDeviceWithNames(req *ApplyDeviceWithNamesArgs) (resp *ApplyDeviceWithNamesResponse, err error) {
	resp = &ApplyDeviceWithNamesResponse{}
	err = c.Request("ApplyDeviceWithNames", req, resp)
	return
}

type ApplyDeviceWithNamesArgs struct {
	DeviceNames ApplyDeviceWithNamesDeviceNameList
	ProductKey  string
}
type ApplyDeviceWithNamesResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	ApplyId      int64
}

type ApplyDeviceWithNamesDeviceNameList []string

func (list *ApplyDeviceWithNamesDeviceNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *IotClient) SaveDeviceProp(req *SaveDevicePropArgs) (resp *SaveDevicePropResponse, err error) {
	resp = &SaveDevicePropResponse{}
	err = c.Request("SaveDeviceProp", req, resp)
	return
}

type SaveDevicePropArgs struct {
	DeviceName string
	ProductKey string
	Props      string
}
type SaveDevicePropResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
}

func (c *IotClient) RegistDevice(req *RegistDeviceArgs) (resp *RegistDeviceResponse, err error) {
	resp = &RegistDeviceResponse{}
	err = c.Request("RegistDevice", req, resp)
	return
}

type RegistDeviceArgs struct {
	DeviceName string
	ProductKey string
}
type RegistDeviceResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	DeviceId     string
	DeviceSecret string
	DeviceStatus string
	DeviceName   string
}

func (c *IotClient) QueryDevice(req *QueryDeviceArgs) (resp *QueryDeviceResponse, err error) {
	resp = &QueryDeviceResponse{}
	err = c.Request("QueryDevice", req, resp)
	return
}

type QueryDeviceDeviceInfo struct {
	DeviceId     string
	DeviceSecret string
	ProductKey   string
	DeviceStatus string
	DeviceName   string
	DeviceType   string
	GmtCreate    string
	GmtModified  string
}
type QueryDeviceArgs struct {
	PageSize    int
	CurrentPage int
	ProductKey  string
}
type QueryDeviceResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	Total        int
	PageSize     int
	PageCount    int
	Page         int
	Data         QueryDeviceDeviceInfoList
}

type QueryDeviceDeviceInfoList []QueryDeviceDeviceInfo

func (list *QueryDeviceDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceDeviceInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *IotClient) UpdateDeviceShadow(req *UpdateDeviceShadowArgs) (resp *UpdateDeviceShadowResponse, err error) {
	resp = &UpdateDeviceShadowResponse{}
	err = c.Request("UpdateDeviceShadow", req, resp)
	return
}

type UpdateDeviceShadowArgs struct {
	ShadowMessage string
	DeviceName    string
	ProductKey    string
}
type UpdateDeviceShadowResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
}

func (c *IotClient) PubBroadcast(req *PubBroadcastArgs) (resp *PubBroadcastResponse, err error) {
	resp = &PubBroadcastResponse{}
	err = c.Request("PubBroadcast", req, resp)
	return
}

type PubBroadcastArgs struct {
	TopicFullName  string
	MessageContent string
	ProductKey     string
}
type PubBroadcastResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
}

func (c *IotClient) UpdateProduct(req *UpdateProductArgs) (resp *UpdateProductResponse, err error) {
	resp = &UpdateProductResponse{}
	err = c.Request("UpdateProduct", req, resp)
	return
}

type UpdateProductArgs struct {
	CatId       int64
	ProductName string
	ExtProps    string
	ProductKey  string
	ProductDesc string
}
type UpdateProductResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
}

func (c *IotClient) BatchGetDeviceState(req *BatchGetDeviceStateArgs) (resp *BatchGetDeviceStateResponse, err error) {
	resp = &BatchGetDeviceStateResponse{}
	err = c.Request("BatchGetDeviceState", req, resp)
	return
}

type BatchGetDeviceStateDeviceStatus struct {
	DeviceId       string
	DeviceName     string
	Status         string
	AsAddress      string
	LastOnlineTime string
}
type BatchGetDeviceStateArgs struct {
	DeviceNames BatchGetDeviceStateDeviceNameList
	ProductKey  string
}
type BatchGetDeviceStateResponse struct {
	RequestId        string
	Success          core.Bool
	ErrorMessage     string
	DeviceStatusList BatchGetDeviceStateDeviceStatusList
}

type BatchGetDeviceStateDeviceNameList []string

func (list *BatchGetDeviceStateDeviceNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type BatchGetDeviceStateDeviceStatusList []BatchGetDeviceStateDeviceStatus

func (list *BatchGetDeviceStateDeviceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BatchGetDeviceStateDeviceStatus)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *IotClient) QueryApplyStatus(req *QueryApplyStatusArgs) (resp *QueryApplyStatusResponse, err error) {
	resp = &QueryApplyStatusResponse{}
	err = c.Request("QueryApplyStatus", req, resp)
	return
}

type QueryApplyStatusArgs struct {
	ApplyId int64
}
type QueryApplyStatusResponse struct {
	RequestId    string
	Success      core.Bool
	ErrorMessage string
	Finish       core.Bool
}
