
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceNetExpireTimeRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
ConnectionString string `position:"Query" name:"ConnectionString"`
ClassicExpiredDays int `position:"Query" name:"ClassicExpiredDays"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r ModifyInstanceNetExpireTimeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceNetExpireTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceNetExpireTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceNetExpireTime", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceNetExpireTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceNetExpireTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceNetExpireTimeResponse struct {
RequestId string
InstanceId string
NetInfoItems ModifyInstanceNetExpireTimeNetInfoItemList
}

type ModifyInstanceNetExpireTimeNetInfoItem struct {
DBInstanceNetType string
Port string
ExpiredTime string
ConnectionString string
IPAddress string
}

                    type ModifyInstanceNetExpireTimeNetInfoItemList []ModifyInstanceNetExpireTimeNetInfoItem

                    func (list *ModifyInstanceNetExpireTimeNetInfoItemList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]ModifyInstanceNetExpireTimeNetInfoItem)
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
                    
