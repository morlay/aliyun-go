package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeMonthlyServiceStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Month                string `position:"Query" name:"Month"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeMonthlyServiceStatusRequest) Invoke(client *sdk.Client) (resp *DescribeMonthlyServiceStatusResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonthlyServiceStatus", "redisa", "")
	resp = &DescribeMonthlyServiceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMonthlyServiceStatusResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalCount       common.Long
	InstanceSLAInfos DescribeMonthlyServiceStatusInstanceSLAInfoList
}

type DescribeMonthlyServiceStatusInstanceSLAInfo struct {
	InstanceId common.String
	UptimePct  common.Float
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
