package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBClusterIPArrayListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBClusterIPArrayListRequest) Invoke(client *sdk.Client) (response *DescribeDBClusterIPArrayListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBClusterIPArrayListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterIPArrayList", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBClusterIPArrayListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBClusterIPArrayListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBClusterIPArrayListResponse struct {
	RequestId string
	Items     DescribeDBClusterIPArrayListDBClusterIPArrayList
}

type DescribeDBClusterIPArrayListDBClusterIPArray struct {
	DBClusterIPArrayName      string
	DBClusterIPArrayAttribute string
	SecurityIPList            string
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
