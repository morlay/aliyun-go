
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeSnapshotsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
SnapshotIds string `position:"Domain" name:"SnapshotIds"`
EndTime string `position:"Query" name:"EndTime"`
BeginTime string `position:"Query" name:"BeginTime"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeSnapshotsRequest) Invoke(client *sdk.Client) (response *DescribeSnapshotsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSnapshotsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeSnapshots", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSnapshotsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeSnapshotsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSnapshotsResponse struct {
RequestId string
Snapshots DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshot struct {
SnapshotId string
SnapshotName string
InstanceId string
CreateTime string
Memory int64
RdbSize int64
Status string
Type string
OssDownloadInPath string
OssDownloadOutPath string
}

                    type DescribeSnapshotsSnapshotList []DescribeSnapshotsSnapshot

                    func (list *DescribeSnapshotsSnapshotList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]DescribeSnapshotsSnapshot)
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
                    
