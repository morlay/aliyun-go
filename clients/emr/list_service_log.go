package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListServiceLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListServiceLogRequest) Invoke(client *sdk.Client) (resp *ListServiceLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListServiceLog", "", "")
	resp = &ListServiceLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListServiceLogResponse struct {
	responses.BaseResponse
	RequestId   string
	LogFileList ListServiceLogLogFileList
}

type ListServiceLogLogFile struct {
	FileName     string
	Size         int64
	HostName     string
	LastModified int64
}

type ListServiceLogLogFileList []ListServiceLogLogFile

func (list *ListServiceLogLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListServiceLogLogFile)
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
