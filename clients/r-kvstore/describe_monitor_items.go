package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMonitorItemsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeMonitorItemsRequest) Invoke(client *sdk.Client) (resp *DescribeMonitorItemsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonitorItems", "redisa", "")
	resp = &DescribeMonitorItemsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMonitorItemsResponse struct {
	responses.BaseResponse
	RequestId    string
	MonitorItems DescribeMonitorItemsKVStoreMonitorItemList
}

type DescribeMonitorItemsKVStoreMonitorItem struct {
	MonitorKey string
	Unit       string
}

type DescribeMonitorItemsKVStoreMonitorItemList []DescribeMonitorItemsKVStoreMonitorItem

func (list *DescribeMonitorItemsKVStoreMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsKVStoreMonitorItem)
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
