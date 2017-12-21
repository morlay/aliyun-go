
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeSecurityIpsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeSecurityIpsRequest) Invoke(client *sdk.Client) (response *DescribeSecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeSecurityIps", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSecurityIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeSecurityIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSecurityIpsResponse struct {
RequestId string
SecurityIpGroups DescribeSecurityIpsSecurityIpGroupList
}

type DescribeSecurityIpsSecurityIpGroup struct {
SecurityIpGroupName string
SecurityIpGroupAttribute string
SecurityIpList string
}

                    type DescribeSecurityIpsSecurityIpGroupList []DescribeSecurityIpsSecurityIpGroup

                    func (list *DescribeSecurityIpsSecurityIpGroupList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeSecurityIpsSecurityIpGroup)
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
                    
