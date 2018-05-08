package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveRecordConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveRecordConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveRecordConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveRecordConfig", "", "")
	resp = &DescribeLiveRecordConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveRecordConfigResponse struct {
	responses.BaseResponse
	RequestId         common.String
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName      common.String
	AppName         common.String
	OssEndpoint     common.String
	OssBucket       common.String
	OssObjectPrefix common.String
	CreateTime      common.String
}

type DescribeLiveRecordConfigLiveAppRecordList []DescribeLiveRecordConfigLiveAppRecord

func (list *DescribeLiveRecordConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigLiveAppRecord)
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
