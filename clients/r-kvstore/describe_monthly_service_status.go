
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeMonthlyServiceStatusRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
Month string `position:"Query" name:"Month"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
InstanceIds string `position:"Query" name:"InstanceIds"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeMonthlyServiceStatusRequest) Invoke(client *sdk.Client) (response *DescribeMonthlyServiceStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMonthlyServiceStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonthlyServiceStatus", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMonthlyServiceStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeMonthlyServiceStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMonthlyServiceStatusResponse struct {
RequestId string
TotalCount int64
InstanceSLAInfos DescribeMonthlyServiceStatusInstanceSLAInfoList
}

type DescribeMonthlyServiceStatusInstanceSLAInfo struct {
InstanceId string
UptimePct float32
}

                    type DescribeMonthlyServiceStatusInstanceSLAInfoList []DescribeMonthlyServiceStatusInstanceSLAInfo

                    func (list *DescribeMonthlyServiceStatusInstanceSLAInfoList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeMonthlyServiceStatusInstanceSLAInfo)
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
                    
