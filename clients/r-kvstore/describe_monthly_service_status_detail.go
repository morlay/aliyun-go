
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeMonthlyServiceStatusDetailRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
Month string `position:"Query" name:"Month"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeMonthlyServiceStatusDetailRequest) Invoke(client *sdk.Client) (response *DescribeMonthlyServiceStatusDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMonthlyServiceStatusDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonthlyServiceStatusDetail", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMonthlyServiceStatusDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeMonthlyServiceStatusDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMonthlyServiceStatusDetailResponse struct {
RequestId string
InstanceId string
UptimePct float32
AffectedInfos DescribeMonthlyServiceStatusDetailAffectedInfoList
}

type DescribeMonthlyServiceStatusDetailAffectedInfo struct {
StartTime string
EndTime string
Description string
}

                    type DescribeMonthlyServiceStatusDetailAffectedInfoList []DescribeMonthlyServiceStatusDetailAffectedInfo

                    func (list *DescribeMonthlyServiceStatusDetailAffectedInfoList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeMonthlyServiceStatusDetailAffectedInfo)
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
                    
