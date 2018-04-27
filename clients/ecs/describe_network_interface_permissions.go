package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNetworkInterfacePermissionsRequest struct {
	requests.RpcRequest
	ResourceOwnerId               int64                                                                `position:"Query" name:"ResourceOwnerId"`
	PageNumber                    int                                                                  `position:"Query" name:"PageNumber"`
	PageSize                      int                                                                  `position:"Query" name:"PageSize"`
	NetworkInterfacePermissionIds *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList `position:"Query" type:"Repeated" name:"NetworkInterfacePermissionId"`
	ResourceOwnerAccount          string                                                               `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string                                                               `position:"Query" name:"OwnerAccount"`
	OwnerId                       int64                                                                `position:"Query" name:"OwnerId"`
	NetworkInterfaceId            string                                                               `position:"Query" name:"NetworkInterfaceId"`
}

func (req *DescribeNetworkInterfacePermissionsRequest) Invoke(client *sdk.Client) (resp *DescribeNetworkInterfacePermissionsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfacePermissions", "ecs", "")
	resp = &DescribeNetworkInterfacePermissionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNetworkInterfacePermissionsResponse struct {
	responses.BaseResponse
	RequestId                   string
	TotalCount                  int
	PageNumber                  int
	PageSize                    int
	NetworkInterfacePermissions DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList
}

type DescribeNetworkInterfacePermissionsNetworkInterfacePermission struct {
	AccountId                    int64
	ServiceName                  string
	NetworkInterfaceId           string
	NetworkInterfacePermissionId string
	Permission                   string
	PermissionState              string
}

type DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList []string

func (list *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList []DescribeNetworkInterfacePermissionsNetworkInterfacePermission

func (list *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacePermissionsNetworkInterfacePermission)
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
