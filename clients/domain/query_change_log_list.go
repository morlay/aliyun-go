package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryChangeLogListRequest struct {
	requests.RpcRequest
	EndDate      int64  `position:"Query" name:"EndDate"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Lang         string `position:"Query" name:"Lang"`
	PageNum      int    `position:"Query" name:"PageNum"`
	StartDate    int64  `position:"Query" name:"StartDate"`
}

func (req *QueryChangeLogListRequest) Invoke(client *sdk.Client) (resp *QueryChangeLogListResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryChangeLogList", "", "")
	resp = &QueryChangeLogListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryChangeLogListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	ResultLimit    bool
	Data           QueryChangeLogListChangeLogList
}

type QueryChangeLogListChangeLog struct {
	DomainName         string
	Result             string
	Operation          string
	OperationIPAddress string
	Details            string
	Time               string
}

type QueryChangeLogListChangeLogList []QueryChangeLogListChangeLog

func (list *QueryChangeLogListChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryChangeLogListChangeLog)
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
