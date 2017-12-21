
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeDBInstanceNetInfoRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceNetInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceNetInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeDBInstanceNetInfo", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceNetInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeDBInstanceNetInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
RequestId string
InstanceNetworkType string
NetInfoItems DescribeDBInstanceNetInfoInstanceNetInfoList
}

type DescribeDBInstanceNetInfoInstanceNetInfo struct {
ConnectionString string
IPAddress string
Port string
VPCId string
VSwitchId string
DBInstanceNetType string
ExpiredTime string
Upgradeable string
}

                    type DescribeDBInstanceNetInfoInstanceNetInfoList []DescribeDBInstanceNetInfoInstanceNetInfo

                    func (list *DescribeDBInstanceNetInfoInstanceNetInfoList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeDBInstanceNetInfoInstanceNetInfo)
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
                    
