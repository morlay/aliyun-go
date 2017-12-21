package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveRecordConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveRecordConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveRecordConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveRecordConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveRecordConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveRecordConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveRecordConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveRecordConfigResponse struct {
	RequestId         string
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName      string
	AppName         string
	OssEndpoint     string
	OssBucket       string
	OssObjectPrefix string
	CreateTime      string
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
