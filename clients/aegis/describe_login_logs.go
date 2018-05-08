package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLoginLogsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	PageSize        int    `position:"Query" name:"PageSize"`
	CurrentPage     int    `position:"Query" name:"CurrentPage"`
}

func (req *DescribeLoginLogsRequest) Invoke(client *sdk.Client) (resp *DescribeLoginLogsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeLoginLogs", "vipaegis", "")
	resp = &DescribeLoginLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoginLogsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PageSize       common.Integer
	CurrentPage    common.Integer
	TotalCount     common.Integer
	HttpStatusCode common.Integer
	LoginLogs      DescribeLoginLogsLoginLogList
}

type DescribeLoginLogsLoginLogList []common.String

func (list *DescribeLoginLogsLoginLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
