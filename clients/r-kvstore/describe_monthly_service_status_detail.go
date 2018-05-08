package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeMonthlyServiceStatusDetailRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	Month                string `position:"Query" name:"Month"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeMonthlyServiceStatusDetailRequest) Invoke(client *sdk.Client) (resp *DescribeMonthlyServiceStatusDetailResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonthlyServiceStatusDetail", "redisa", "")
	resp = &DescribeMonthlyServiceStatusDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMonthlyServiceStatusDetailResponse struct {
	responses.BaseResponse
	RequestId     common.String
	InstanceId    common.String
	UptimePct     common.Float
	AffectedInfos DescribeMonthlyServiceStatusDetailAffectedInfoList
}

type DescribeMonthlyServiceStatusDetailAffectedInfo struct {
	StartTime   common.String
	EndTime     common.String
	Description common.String
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
