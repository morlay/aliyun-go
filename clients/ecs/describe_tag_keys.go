package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTagKeysRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeTagKeysRequest) Invoke(client *sdk.Client) (resp *DescribeTagKeysResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTagKeys", "ecs", "")
	resp = &DescribeTagKeysResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTagKeysResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageSize   common.Integer
	PageNumber common.Integer
	TotalCount common.Integer
	TagKeys    DescribeTagKeysTagKeyList
}

type DescribeTagKeysTagKeyList []common.String

func (list *DescribeTagKeysTagKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
