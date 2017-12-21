
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeMonitorItemsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeMonitorItemsRequest) Invoke(client *sdk.Client) (response *DescribeMonitorItemsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMonitorItemsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonitorItems", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMonitorItemsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeMonitorItemsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMonitorItemsResponse struct {
RequestId string
MonitorItems DescribeMonitorItemsKVStoreMonitorItemList
}

type DescribeMonitorItemsKVStoreMonitorItem struct {
MonitorKey string
Unit string
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
                    
