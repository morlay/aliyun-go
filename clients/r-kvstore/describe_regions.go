
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeRegionsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeRegionsRequest) Invoke(client *sdk.Client) (response *DescribeRegionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRegionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeRegions", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRegionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeRegionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRegionsResponse struct {
RequestId string
RegionIds DescribeRegionsKVStoreRegionList
}

type DescribeRegionsKVStoreRegion struct {
RegionId string
ZoneIds string
LocalName string
}

                    type DescribeRegionsKVStoreRegionList []DescribeRegionsKVStoreRegion

                    func (list *DescribeRegionsKVStoreRegionList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeRegionsKVStoreRegion)
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
                    
