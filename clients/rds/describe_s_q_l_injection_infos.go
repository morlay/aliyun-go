package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLInjectionInfosRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSQLInjectionInfosRequest) Invoke(client *sdk.Client) (resp *DescribeSQLInjectionInfosResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLInjectionInfos", "rds", "")
	resp = &DescribeSQLInjectionInfosResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLInjectionInfosResponse struct {
	responses.BaseResponse
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLInjectionInfosSQLInjectionInfoList
}

type DescribeSQLInjectionInfosSQLInjectionInfo struct {
	DBName         string
	SQLText        string
	LatencyTime    string
	HostAddress    string
	ExecuteTime    string
	AccountName    string
	EffectRowCount string
}

type DescribeSQLInjectionInfosSQLInjectionInfoList []DescribeSQLInjectionInfosSQLInjectionInfo

func (list *DescribeSQLInjectionInfosSQLInjectionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLInjectionInfosSQLInjectionInfo)
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
