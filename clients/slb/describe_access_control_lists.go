package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAccessControlListsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AclName              string `position:"Query" name:"AclName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeAccessControlListsRequest) Invoke(client *sdk.Client) (resp *DescribeAccessControlListsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeAccessControlLists", "slb", "")
	resp = &DescribeAccessControlListsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessControlListsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Acls      DescribeAccessControlListsAclList
}

type DescribeAccessControlListsAcl struct {
	AclId   common.String
	AclName common.String
}

type DescribeAccessControlListsAclList []DescribeAccessControlListsAcl

func (list *DescribeAccessControlListsAclList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListsAcl)
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
