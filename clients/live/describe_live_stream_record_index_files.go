package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRecordIndexFilesRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	Order         string `position:"Query" name:"Order"`
}

func (req *DescribeLiveStreamRecordIndexFilesRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamRecordIndexFiles", "", "")
	resp = &DescribeLiveStreamRecordIndexFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	responses.BaseResponse
	RequestId           string
	PageNum             int
	PageSize            int
	Order               string
	TotalNum            int
	TotalPage           int
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId    string
	RecordUrl   string
	DomainName  string
	AppName     string
	StreamName  string
	OssBucket   string
	OssEndpoint string
	OssObject   string
	StartTime   string
	EndTime     string
	Duration    float32
	Height      int
	Width       int
	CreateTime  string
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfoList []DescribeLiveStreamRecordIndexFilesRecordIndexInfo

func (list *DescribeLiveStreamRecordIndexFilesRecordIndexInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordIndexFilesRecordIndexInfo)
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
