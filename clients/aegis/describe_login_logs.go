package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	PageSize       int
	CurrentPage    int
	TotalCount     int
	HttpStatusCode int
	LoginLogs      DescribeLoginLogsLoginLogList
}

type DescribeLoginLogsLoginLogList []string

func (list *DescribeLoginLogsLoginLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
