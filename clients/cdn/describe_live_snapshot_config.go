package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveSnapshotConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	Order         string `position:"Query" name:"Order"`
}

func (r DescribeLiveSnapshotConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveSnapshotConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveSnapshotConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveSnapshotConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveSnapshotConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveSnapshotConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveSnapshotConfigResponse struct {
	RequestId                    string
	PageNum                      int
	PageSize                     int
	Order                        string
	TotalNum                     int
	TotalPage                    int
	LiveStreamSnapshotConfigList DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfig struct {
	DomainName         string
	AppName            string
	TimeInterval       int
	OssEndpoint        string
	OssBucket          string
	OverwriteOssObject string
	SequenceOssObject  string
	CreateTime         string
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList []DescribeLiveSnapshotConfigLiveStreamSnapshotConfig

func (list *DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotConfigLiveStreamSnapshotConfig)
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
