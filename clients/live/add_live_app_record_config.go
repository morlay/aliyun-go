package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveAppRecordConfigRequest struct {
	OssBucket     string                                  `position:"Query" name:"OssBucket"`
	AppName       string                                  `position:"Query" name:"AppName"`
	SecurityToken string                                  `position:"Query" name:"SecurityToken"`
	RecordFormats *AddLiveAppRecordConfigRecordFormatList `position:"Query" type:"Repeated" name:"RecordFormat"`
	DomainName    string                                  `position:"Query" name:"DomainName"`
	OssEndpoint   string                                  `position:"Query" name:"OssEndpoint"`
	OwnerId       int64                                   `position:"Query" name:"OwnerId"`
}

func (r AddLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (response *AddLiveAppRecordConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveAppRecordConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveAppRecordConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveAppRecordConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveAppRecordConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveAppRecordConfigRecordFormat struct {
	Format               string `name:"Format"`
	OssObjectPrefix      string `name:"OssObjectPrefix"`
	SliceOssObjectPrefix string `name:"SliceOssObjectPrefix"`
	CycleDuration        int    `name:"CycleDuration"`
}

type AddLiveAppRecordConfigResponse struct {
	RequestId string
}

type AddLiveAppRecordConfigRecordFormatList []AddLiveAppRecordConfigRecordFormat

func (list *AddLiveAppRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddLiveAppRecordConfigRecordFormat)
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
