package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveSnapshotDetectPornConfigRequest struct {
	requests.RpcRequest
	OssBucket     string                                       `position:"Query" name:"OssBucket"`
	AppName       string                                       `position:"Query" name:"AppName"`
	SecurityToken string                                       `position:"Query" name:"SecurityToken"`
	DomainName    string                                       `position:"Query" name:"DomainName"`
	OssEndpoint   string                                       `position:"Query" name:"OssEndpoint"`
	Interval      int                                          `position:"Query" name:"Interval"`
	OwnerId       int64                                        `position:"Query" name:"OwnerId"`
	OssObject     string                                       `position:"Query" name:"OssObject"`
	Scenes        *UpdateLiveSnapshotDetectPornConfigSceneList `position:"Query" type:"Repeated" name:"Scene"`
}

func (req *UpdateLiveSnapshotDetectPornConfigRequest) Invoke(client *sdk.Client) (resp *UpdateLiveSnapshotDetectPornConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveSnapshotDetectPornConfig", "", "")
	resp = &UpdateLiveSnapshotDetectPornConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLiveSnapshotDetectPornConfigResponse struct {
	responses.BaseResponse
	RequestId string
}

type UpdateLiveSnapshotDetectPornConfigSceneList []string

func (list *UpdateLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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
