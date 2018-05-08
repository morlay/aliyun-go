package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAccessControlListAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	AclId                string `position:"Query" name:"AclId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AclEntryComment      string `position:"Query" name:"AclEntryComment"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeAccessControlListAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeAccessControlListAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeAccessControlListAttribute", "slb", "")
	resp = &DescribeAccessControlListAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessControlListAttributeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	AclId            common.String
	AclName          common.String
	AclEntrys        DescribeAccessControlListAttributeAclEntryList
	RelatedListeners DescribeAccessControlListAttributeRelatedListenerList
}

type DescribeAccessControlListAttributeAclEntry struct {
	AclEntryIP      common.String
	AclEntryComment common.String
}

type DescribeAccessControlListAttributeRelatedListener struct {
	LoadBalancerId common.String
	ListenerPort   common.Integer
	AclType        common.String
	Protocol       common.String
}

type DescribeAccessControlListAttributeAclEntryList []DescribeAccessControlListAttributeAclEntry

func (list *DescribeAccessControlListAttributeAclEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListAttributeAclEntry)
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

type DescribeAccessControlListAttributeRelatedListenerList []DescribeAccessControlListAttributeRelatedListener

func (list *DescribeAccessControlListAttributeRelatedListenerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessControlListAttributeRelatedListener)
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
