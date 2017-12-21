package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImageSharePermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeImageSharePermissionRequest) Invoke(client *sdk.Client) (response *DescribeImageSharePermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeImageSharePermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageSharePermission", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeImageSharePermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeImageSharePermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeImageSharePermissionResponse struct {
	RequestId   string
	RegionId    string
	TotalCount  int
	PageNumber  int
	PageSize    int
	ImageId     string
	ShareGroups DescribeImageSharePermissionShareGroupList
	Accounts    DescribeImageSharePermissionAccountList
}

type DescribeImageSharePermissionShareGroup struct {
	Group string
}

type DescribeImageSharePermissionAccount struct {
	AliyunId string
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
