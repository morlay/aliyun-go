package hpc

import (
	"encoding/json"
)

func (c *HpcClient) AuthorizeSecurityGroup(req *AuthorizeSecurityGroupArgs) (resp *AuthorizeSecurityGroupResponse, err error) {
	resp = &AuthorizeSecurityGroupResponse{}
	err = c.Request("AuthorizeSecurityGroup", req, resp)
	return
}

type AuthorizeSecurityGroupArgs struct {
	TOKEN    string
	SourceIp string
	Policy   string
	Priority string
	NicType  string
}
type AuthorizeSecurityGroupResponse struct {
	RequestId string
}

func (c *HpcClient) RevokeSecurityGroup(req *RevokeSecurityGroupArgs) (resp *RevokeSecurityGroupResponse, err error) {
	resp = &RevokeSecurityGroupResponse{}
	err = c.Request("RevokeSecurityGroup", req, resp)
	return
}

type RevokeSecurityGroupArgs struct {
	TOKEN    string
	SourceIp string
	Policy   string
	Priority string
	NicType  string
}
type RevokeSecurityGroupResponse struct {
	RequestId string
}

func (c *HpcClient) DescribeInstancesInSecurityGroup(req *DescribeInstancesInSecurityGroupArgs) (resp *DescribeInstancesInSecurityGroupResponse, err error) {
	resp = &DescribeInstancesInSecurityGroupResponse{}
	err = c.Request("DescribeInstancesInSecurityGroup", req, resp)
	return
}

type DescribeInstancesInSecurityGroupRecord struct {
	RegionId       string
	InstanceId     string
	InnerIpAddress string
}
type DescribeInstancesInSecurityGroupArgs struct {
	TOKEN string
}
type DescribeInstancesInSecurityGroupResponse struct {
	RequestId string
	Records   DescribeInstancesInSecurityGroupRecordList
}

type DescribeInstancesInSecurityGroupRecordList []DescribeInstancesInSecurityGroupRecord

func (list *DescribeInstancesInSecurityGroupRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInSecurityGroupRecord)
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

func (c *HpcClient) DescribeInstances(req *DescribeInstancesArgs) (resp *DescribeInstancesResponse, err error) {
	resp = &DescribeInstancesResponse{}
	err = c.Request("DescribeInstances", req, resp)
	return
}

type DescribeInstancesInstance struct {
	InstanceId                string
	RegionId                  string
	InstanceType              DescribeInstancesInstanceType
	PackageId                 DescribeInstancesPackageId
	Status                    DescribeInstancesStatus
	InnerIpAddress            string
	JumpserverStatus          DescribeInstancesJumpserverStatus
	JumpserverInnerIpAddress  string
	JumpServerPublicIpAddress string
	CreationTime              string
	ExpireTime                string
}

type DescribeInstancesInstanceType struct {
	StringValue string
}

type DescribeInstancesPackageId struct {
	StringValue string
}

type DescribeInstancesStatus struct {
	StringValue string
}

type DescribeInstancesJumpserverStatus struct {
	StringValue string
}
type DescribeInstancesArgs struct {
	TOKEN        string
	InstanceId   string
	InstanceType string
}
type DescribeInstancesResponse struct {
	RequestId string
	Instances DescribeInstancesInstanceList
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

func (c *HpcClient) RebootJumpserver(req *RebootJumpserverArgs) (resp *RebootJumpserverResponse, err error) {
	resp = &RebootJumpserverResponse{}
	err = c.Request("RebootJumpserver", req, resp)
	return
}

type RebootJumpserverArgs struct {
	TOKEN      string
	InstanceId string
	Force      int
}
type RebootJumpserverResponse struct {
	RequestId string
}

func (c *HpcClient) StopJumpserver(req *StopJumpserverArgs) (resp *StopJumpserverResponse, err error) {
	resp = &StopJumpserverResponse{}
	err = c.Request("StopJumpserver", req, resp)
	return
}

type StopJumpserverArgs struct {
	TOKEN      string
	InstanceId string
	Force      int
}
type StopJumpserverResponse struct {
	RequestId string
}

func (c *HpcClient) RebootInstance(req *RebootInstanceArgs) (resp *RebootInstanceResponse, err error) {
	resp = &RebootInstanceResponse{}
	err = c.Request("RebootInstance", req, resp)
	return
}

type RebootInstanceArgs struct {
	TOKEN      string
	InstanceId string
}
type RebootInstanceResponse struct {
	RequestId string
}

func (c *HpcClient) ModifyInstancePassword(req *ModifyInstancePasswordArgs) (resp *ModifyInstancePasswordResponse, err error) {
	resp = &ModifyInstancePasswordResponse{}
	err = c.Request("ModifyInstancePassword", req, resp)
	return
}

type ModifyInstancePasswordArgs struct {
	TOKEN       string
	InstanceId  string
	NewPassword string
}
type ModifyInstancePasswordResponse struct {
	RequestId string
}

func (c *HpcClient) ModifyJumpserverPassword(req *ModifyJumpserverPasswordArgs) (resp *ModifyJumpserverPasswordResponse, err error) {
	resp = &ModifyJumpserverPasswordResponse{}
	err = c.Request("ModifyJumpserverPassword", req, resp)
	return
}

type ModifyJumpserverPasswordArgs struct {
	TOKEN       string
	InstanceId  string
	NewPassword string
}
type ModifyJumpserverPasswordResponse struct {
	RequestId string
}

func (c *HpcClient) DescribeSecurityGroupAttribute(req *DescribeSecurityGroupAttributeArgs) (resp *DescribeSecurityGroupAttributeResponse, err error) {
	resp = &DescribeSecurityGroupAttributeResponse{}
	err = c.Request("DescribeSecurityGroupAttribute", req, resp)
	return
}

type DescribeSecurityGroupAttributeRecords struct {
	RegionId    string
	Permissions DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	SourceIp string
	Policy   DescribeSecurityGroupAttributePolicy
	NicType  DescribeSecurityGroupAttributeNicType
	Priority string
	Time     string
}

type DescribeSecurityGroupAttributePolicy struct {
	StringValue string
}

type DescribeSecurityGroupAttributeNicType struct {
	StringValue string
}
type DescribeSecurityGroupAttributeArgs struct {
	TOKEN   string
	NicType string
}
type DescribeSecurityGroupAttributeResponse struct {
	Records DescribeSecurityGroupAttributeRecords
}

type DescribeSecurityGroupAttributePermissionList []DescribeSecurityGroupAttributePermission

func (list *DescribeSecurityGroupAttributePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupAttributePermission)
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

func (c *HpcClient) StartJumpserver(req *StartJumpserverArgs) (resp *StartJumpserverResponse, err error) {
	resp = &StartJumpserverResponse{}
	err = c.Request("StartJumpserver", req, resp)
	return
}

type StartJumpserverArgs struct {
	TOKEN      string
	InstanceId string
	Force      int
}
type StartJumpserverResponse struct {
	RequestId string
}
