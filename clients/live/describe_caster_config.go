package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCasterConfigRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCasterConfigRequest) Invoke(client *sdk.Client) (resp *DescribeCasterConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterConfig", "live", "")
	resp = &DescribeCasterConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterConfigResponse struct {
	responses.BaseResponse
	RequestId        common.String
	CasterId         common.String
	CasterName       common.String
	DomainName       common.String
	Delay            common.Float
	UrgentMaterialId common.String
	SideOutputUrl    common.String
	CallbackUrl      common.String
	ProgramName      common.String
	ProgramEffect    common.Integer
	TranscodeConfig  DescribeCasterConfigTranscodeConfig
	RecordConfig     DescribeCasterConfigRecordConfig
}

type DescribeCasterConfigTranscodeConfig struct {
	CasterTemplate  common.String
	LiveTemplateIds DescribeCasterConfigLiveTemplateIdList
}

type DescribeCasterConfigRecordConfig struct {
	OssEndpoint  common.String
	OssBucket    common.String
	RecordFormat DescribeCasterConfigRecordFormatItemList
}

type DescribeCasterConfigRecordFormatItem struct {
	Format               common.String
	OssObjectPrefix      common.String
	SliceOssObjectPrefix common.String
	CycleDuration        common.Integer
}

type DescribeCasterConfigLiveTemplateIdList []common.String

func (list *DescribeCasterConfigLiveTemplateIdList) UnmarshalJSON(data []byte) error {
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
