package waf

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *WafClient) RenewInstance(req *RenewInstanceArgs) (resp *RenewInstanceResponse, err error) {
	resp = &RenewInstanceResponse{}
	err = c.Request("RenewInstance", req, resp)
	return
}

type RenewInstanceArgs struct {
	OwnerId      int64
	ClientToken  string
	InstanceId   string
	Duration     int
	PricingCycle string
}
type RenewInstanceResponse struct {
	OrderId   string
	RequestId string
}

func (c *WafClient) CreateInstance(req *CreateInstanceArgs) (resp *CreateInstanceResponse, err error) {
	resp = &CreateInstanceResponse{}
	err = c.Request("CreateInstance", req, resp)
	return
}

type CreateInstanceArgs struct {
	OwnerId           int64
	ClientToken       string
	Duration          int
	PricingCycle      string
	PackageCode       string
	ExtDomainPackage  int
	ExtBandwidth      int
	IsAutoRenew       core.Bool
	AutoRenewDuration int
}
type CreateInstanceResponse struct {
	OrderId    string
	InstanceId string
	RequestId  string
}

func (c *WafClient) ReleaseInstance(req *ReleaseInstanceArgs) (resp *ReleaseInstanceResponse, err error) {
	resp = &ReleaseInstanceResponse{}
	err = c.Request("ReleaseInstance", req, resp)
	return
}

type ReleaseInstanceArgs struct {
	OwnerId    int64
	InstanceId string
}
type ReleaseInstanceResponse struct {
	RequestId string
}

func (c *WafClient) UpgradeInstance(req *UpgradeInstanceArgs) (resp *UpgradeInstanceResponse, err error) {
	resp = &UpgradeInstanceResponse{}
	err = c.Request("UpgradeInstance", req, resp)
	return
}

type UpgradeInstanceArgs struct {
	OwnerId          int64
	ClientToken      string
	InstanceId       string
	PackageCode      string
	ExtDomainPackage int
	ExtBandwidth     int
}
type UpgradeInstanceResponse struct {
	OrderId   string
	RequestId string
}
