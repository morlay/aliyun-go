package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostInnerIp     string `position:"Query" name:"HostInnerIp"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	FromTimestamp   int    `position:"Query" name:"FromTimestamp"`
	Offset          int    `position:"Query" name:"Offset"`
	Line            int    `position:"Query" name:"Line"`
	ToTimestamp     int    `position:"Query" name:"ToTimestamp"`
	SlsQueryString  string `position:"Query" name:"SlsQueryString"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Reverse         string `position:"Query" name:"Reverse"`
}

func (req *SearchLogRequest) Invoke(client *sdk.Client) (resp *SearchLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SearchLog", "", "")
	resp = &SearchLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchLogResponse struct {
	responses.BaseResponse
	RequestId      string
	Completed      bool
	SlsLogItemList SearchLogSlsLogItemList
}

type SearchLogSlsLogItem struct {
	Timestamp int
	SourceIp  string
	HostName  string
	Path      string
	Content   string
}

type SearchLogSlsLogItemList []SearchLogSlsLogItem

func (list *SearchLogSlsLogItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchLogSlsLogItem)
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
