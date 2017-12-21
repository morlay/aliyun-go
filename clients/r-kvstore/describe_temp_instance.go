
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeTempInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeTempInstanceRequest) Invoke(client *sdk.Client) (response *DescribeTempInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTempInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeTempInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTempInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeTempInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTempInstanceResponse struct {
RequestId string
TempInstances DescribeTempInstanceTempInstanceList
}

type DescribeTempInstanceTempInstance struct {
InstanceId string
TempInstanceId string
SnapshotId string
CreateTime string
Domain string
Status string
Memory int64
ExpireTime string
}

                    type DescribeTempInstanceTempInstanceList []DescribeTempInstanceTempInstance

                    func (list *DescribeTempInstanceTempInstanceList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeTempInstanceTempInstance)
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
                    
