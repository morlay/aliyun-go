package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveRecordVodConfigsRequest struct {
	requests.RpcRequest
	AppName    string `position:"Query" name:"AppName"`
	DomainName string `position:"Query" name:"DomainName"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	PageNum    int64  `position:"Query" name:"PageNum"`
	StreamName string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveRecordVodConfigsRequest) Invoke(client *sdk.Client) (resp *DescribeLiveRecordVodConfigsResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordVodConfigs", "live", "")
	resp = &DescribeLiveRecordVodConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveRecordVodConfigsResponse struct {
	responses.BaseResponse
	RequestId            string
	PageNum              int
	PageSize             int
	Total                string
	LiveRecordVodConfigs DescribeLiveRecordVodConfigsLiveRecordVodConfigList
}

type DescribeLiveRecordVodConfigsLiveRecordVodConfig struct {
	CreateTime                 string
	DomainName                 string
	AppName                    string
	StreamName                 string
	VodTranscodeGroupId        string
	CycleDuration              int
	AutoCompose                string
	ComposeVodTranscodeGroupId string
}

type DescribeLiveRecordVodConfigsLiveRecordVodConfigList []DescribeLiveRecordVodConfigsLiveRecordVodConfig

func (list *DescribeLiveRecordVodConfigsLiveRecordVodConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordVodConfigsLiveRecordVodConfig)
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
