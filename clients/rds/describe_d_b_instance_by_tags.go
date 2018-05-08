package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceByTagsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *DescribeDBInstanceByTagsRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceByTagsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceByTags", "rds", "")
	resp = &DescribeDBInstanceByTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceByTagsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	TotalRecordCount common.Integer
	Items            DescribeDBInstanceByTagsDBInstanceTagList
}

type DescribeDBInstanceByTagsDBInstanceTag struct {
	DBInstanceId common.String
	Tags         DescribeDBInstanceByTagsTagList
}

type DescribeDBInstanceByTagsTag struct {
	TagKey   common.String
	TagValue common.String
}

type DescribeDBInstanceByTagsDBInstanceTagList []DescribeDBInstanceByTagsDBInstanceTag

func (list *DescribeDBInstanceByTagsDBInstanceTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsDBInstanceTag)
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

type DescribeDBInstanceByTagsTagList []DescribeDBInstanceByTagsTag

func (list *DescribeDBInstanceByTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsTag)
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
