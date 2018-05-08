package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeImageSharePermissionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeImageSharePermissionRequest) Invoke(client *sdk.Client) (resp *DescribeImageSharePermissionResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageSharePermission", "ecs", "")
	resp = &DescribeImageSharePermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImageSharePermissionResponse struct {
	responses.BaseResponse
	RequestId   common.String
	RegionId    common.String
	TotalCount  common.Integer
	PageNumber  common.Integer
	PageSize    common.Integer
	ImageId     common.String
	ShareGroups DescribeImageSharePermissionShareGroupList
	Accounts    DescribeImageSharePermissionAccountList
}

type DescribeImageSharePermissionShareGroup struct {
	Group common.String
}

type DescribeImageSharePermissionAccount struct {
	AliyunId common.String
}

type DescribeImageSharePermissionShareGroupList []DescribeImageSharePermissionShareGroup

func (list *DescribeImageSharePermissionShareGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionShareGroup)
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

type DescribeImageSharePermissionAccountList []DescribeImageSharePermissionAccount

func (list *DescribeImageSharePermissionAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionAccount)
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
