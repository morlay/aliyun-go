package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	PageNum              common.Integer
	PageSize             common.Integer
	Total                common.String
	LiveRecordVodConfigs DescribeLiveRecordVodConfigsLiveRecordVodConfigList
}

type DescribeLiveRecordVodConfigsLiveRecordVodConfig struct {
	CreateTime                 common.String
	DomainName                 common.String
	AppName                    common.String
	StreamName                 common.String
	VodTranscodeGroupId        common.String
	CycleDuration              common.Integer
	AutoCompose                common.String
	ComposeVodTranscodeGroupId common.String
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
