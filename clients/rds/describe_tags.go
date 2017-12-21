package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTagsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeTagsRequest) Invoke(client *sdk.Client) (response *DescribeTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeTags", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTagsResponse struct {
	RequestId string
	Items     DescribeTagsTagInfosList
}

type DescribeTagsTagInfos struct {
	TagKey        string
	TagValue      string
	DBInstanceIds DescribeTagsDBInstanceIdList
}

type DescribeTagsTagInfosList []DescribeTagsTagInfos

func (list *DescribeTagsTagInfosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagInfos)
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

type DescribeTagsDBInstanceIdList []string

func (list *DescribeTagsDBInstanceIdList) UnmarshalJSON(data []byte) error {
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
