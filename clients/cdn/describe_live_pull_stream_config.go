package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLivePullStreamConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLivePullStreamConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLivePullStreamConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLivePullStreamConfig", "", "")
	resp = &DescribeLivePullStreamConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLivePullStreamConfigResponse struct {
	responses.BaseResponse
	RequestId         common.String
	LiveAppRecordList DescribeLivePullStreamConfigLiveAppRecordList
}

type DescribeLivePullStreamConfigLiveAppRecord struct {
	DomainName common.String
	StreamName common.String
	SourceUrl  common.String
	StartTime  common.String
	EndTime    common.String
}

type DescribeLivePullStreamConfigLiveAppRecordList []DescribeLivePullStreamConfigLiveAppRecord

func (list *DescribeLivePullStreamConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLivePullStreamConfigLiveAppRecord)
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
