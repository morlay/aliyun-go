package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveSnapshotConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	Order         string `position:"Query" name:"Order"`
}

func (req *DescribeLiveSnapshotConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveSnapshotConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveSnapshotConfig", "", "")
	resp = &DescribeLiveSnapshotConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveSnapshotConfigResponse struct {
	responses.BaseResponse
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
