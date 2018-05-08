package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBClusterIPArrayListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBClusterIPArrayListRequest) Invoke(client *sdk.Client) (resp *DescribeDBClusterIPArrayListResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterIPArrayList", "polardb", "")
	resp = &DescribeDBClusterIPArrayListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBClusterIPArrayListResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeDBClusterIPArrayListDBClusterIPArrayList
}

type DescribeDBClusterIPArrayListDBClusterIPArray struct {
	DBClusterIPArrayName      common.String
	DBClusterIPArrayAttribute common.String
	SecurityIPList            common.String
}

type DescribeDBClusterIPArrayListDBClusterIPArrayList []DescribeDBClusterIPArrayListDBClusterIPArray

func (list *DescribeDBClusterIPArrayListDBClusterIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterIPArrayListDBClusterIPArray)
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
