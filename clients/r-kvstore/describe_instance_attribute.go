
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeInstanceAttributeRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceAttributeRequest) Invoke(client *sdk.Client) (response *DescribeInstanceAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceAttribute", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeInstanceAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceAttributeResponse struct {
RequestId string
Instances DescribeInstanceAttributeDBInstanceAttributeList
}

type DescribeInstanceAttributeDBInstanceAttribute struct {
InstanceId string
InstanceName string
ConnectionDomain string
Port int64
InstanceStatus string
RegionId string
Capacity int64
InstanceClass string
QPS int64
Bandwidth int64
Connections int64
ZoneId string
Config string
ChargeType string
NodeType string
NetworkType string
VpcId string
VSwitchId string
PrivateIp string
CreateTime string
EndTime string
HasRenewChangeOrder string
IsRds bool
Engine string
EngineVersion string
MaintainStartTime string
MaintainEndTime string
AvailabilityValue string
SecurityIPList string
InstanceType string
ArchitectureType string
NodeType1 string
PackageType string
}

                    type DescribeInstanceAttributeDBInstanceAttributeList []DescribeInstanceAttributeDBInstanceAttribute

                    func (list *DescribeInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeInstanceAttributeDBInstanceAttribute)
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
                    
