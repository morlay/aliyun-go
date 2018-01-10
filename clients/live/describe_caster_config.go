package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DescribeCasterConfigRequest) Invoke(client *sdk.Client) (resp *DescribeCasterConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterConfig", "", "")
	resp = &DescribeCasterConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterConfigResponse struct {
	responses.BaseResponse
	RequestId        string
	CasterId         string
	CasterName       string
	DomainName       string
	Delay            float32
	UrgentMaterialId string
	SideOutputUrl    string
	CallbackUrl      string
	TranscodeConfig  DescribeCasterConfigTranscodeConfig
	RecordConfig     DescribeCasterConfigRecordConfig
}

type DescribeCasterConfigTranscodeConfig struct {
	CasterTemplate  string
	LiveTemplateIds DescribeCasterConfigLiveTemplateIdList
}

type DescribeCasterConfigRecordConfig struct {
	OssEndpoint  string
	OssBucket    string
	RecordFormat DescribeCasterConfigRecordFormatItemList
}

type DescribeCasterConfigRecordFormatItem struct {
	Format               string
	OssObjectPrefix      string
	SliceOssObjectPrefix string
	CycleDuration        int
}

type DescribeCasterConfigLiveTemplateIdList []string

func (list *DescribeCasterConfigLiveTemplateIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterConfigRecordFormatItemList []DescribeCasterConfigRecordFormatItem

func (list *DescribeCasterConfigRecordFormatItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterConfigRecordFormatItem)
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
