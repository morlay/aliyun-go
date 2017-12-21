package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveSnapshotDetectPornConfigRequest struct {
	OssBucket     string                                    `position:"Query" name:"OssBucket"`
	AppName       string                                    `position:"Query" name:"AppName"`
	SecurityToken string                                    `position:"Query" name:"SecurityToken"`
	DomainName    string                                    `position:"Query" name:"DomainName"`
	OssEndpoint   string                                    `position:"Query" name:"OssEndpoint"`
	Interval      int                                       `position:"Query" name:"Interval"`
	OwnerId       int64                                     `position:"Query" name:"OwnerId"`
	OssObject     string                                    `position:"Query" name:"OssObject"`
	Scenes        *AddLiveSnapshotDetectPornConfigSceneList `position:"Query" type:"Repeated" name:"Scene"`
}

func (r AddLiveSnapshotDetectPornConfigRequest) Invoke(client *sdk.Client) (response *AddLiveSnapshotDetectPornConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveSnapshotDetectPornConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveSnapshotDetectPornConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveSnapshotDetectPornConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveSnapshotDetectPornConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveSnapshotDetectPornConfigResponse struct {
	RequestId string
}

type AddLiveSnapshotDetectPornConfigSceneList []string

func (list *AddLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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
