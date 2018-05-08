package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicaInitializeProgressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeReplicaInitializeProgressRequest) Invoke(client *sdk.Client) (resp *DescribeReplicaInitializeProgressResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeReplicaInitializeProgress", "redisa", "")
	resp = &DescribeReplicaInitializeProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicaInitializeProgressResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeReplicaInitializeProgressItemsItemList
}

type DescribeReplicaInitializeProgressItemsItem struct {
	ReplicaId   common.String
	Status      common.String
	Progress    common.String
	FinishTime  common.String
	CurrentStep common.String
}

type DescribeReplicaInitializeProgressItemsItemList []DescribeReplicaInitializeProgressItemsItem

func (list *DescribeReplicaInitializeProgressItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaInitializeProgressItemsItem)
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
