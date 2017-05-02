// Package implements OCCM Working Environments API (AWS-HA)
package awsha

// TODO: determine how to improve reuse of methods between VSA and HA APIs
// TODO: consider moving logic to base workenv API and determine the best
// way to simulate inheritance and provide only the endpoint URI

import (
  "fmt"

  "github.com/candidpartners/occm-sdk-go/api/workenv"
  "github.com/candidpartners/occm-sdk-go/api/workenv/vsa"
  "github.com/candidpartners/occm-sdk-go/api/client"
	"github.com/pkg/errors"
)

// VSA Working environment API
type AWSHAWorkingEnvironmentAPI struct {
	*client.Client
}

// New creates a new OCCM VSA Working Environment API client
func New(context *client.Context) (*AWSHAWorkingEnvironmentAPI, error) {
  c, err := client.New(context)
  if err != nil {
    return nil, errors.Wrap(err, client.ErrClientCreationFailed)
  }

	api := &AWSHAWorkingEnvironmentAPI{
		Client: c,
	}

	return api, nil
}

// GetAggregates retrieves a list of all aggregates for a given working environment
func (api *AWSHAWorkingEnvironmentAPI) GetAggregates(workenvId string) ([]workenv.AggregateResponse, error) {
  if workenvId == "" {
		return nil, errors.New(workenv.ErrInvalidWorkenvId)
	}

  data, err := api.Client.Invoke("GET", "/aws/ha/aggregates",
    map[string]string{
      "workingEnvironmentId": workenvId,
    },
    nil,
  )
  if err != nil {
		return nil, errors.Wrap(err, client.ErrInvalidRequest)
	}

  result, err := workenv.AggregateResponseListFromJSON(data);
  if err != nil {
    return nil, errors.Wrap(err, client.ErrJSONConversion)
  }

  return result, nil
}

// GetVolumes retrieves a list of all volumes for a given working environment
func (api *AWSHAWorkingEnvironmentAPI) GetVolumes(workenvId string) ([]workenv.VolumeResponse, error) {
  if workenvId == "" {
		return nil, errors.New(workenv.ErrInvalidWorkenvId)
	}

  data, err := api.Client.Invoke("GET", "/aws/ha/volumes",
    map[string]string{
      "workingEnvironmentId": workenvId,
    },
    nil,
  )
  if err != nil {
		return nil, errors.Wrap(err, client.ErrInvalidRequest)
	}

  result, err := workenv.VolumeResponseListFromJSON(data);
  if err != nil {
    return nil, errors.Wrap(err, client.ErrJSONConversion)
  }

  return result, nil
}

// QuoteVolume quotes a volume for the given request
func (api *AWSHAWorkingEnvironmentAPI) QuoteVolume(request *vsa.VSAVolumeQuoteRequest) (*vsa.VSAVolumeQuoteResponse, error) {
  if request == nil {
		return nil, errors.New(workenv.ErrInvalidVolumeQuoteRequest)
	}

  data, err := api.Client.Invoke("POST", "/aws/ha/volumes/quote", nil, request)
  if err != nil {
		return nil, errors.Wrap(err, client.ErrInvalidRequest)
	}

  result, err := vsa.VolumeQuoteResponseFromJSON(data);
  if err != nil {
    return nil, errors.Wrap(err, client.ErrJSONConversion)
  }

  return result, nil
}

// CreateVolume creates a volume for the given request
func (api *AWSHAWorkingEnvironmentAPI) CreateVolume(createAggregateIfNotFound bool, request *vsa.VSAVolumeCreateRequest) error {
  if request == nil {
		return errors.New(workenv.ErrInvalidVolumeCreationRequest)
	}

  _, err := api.Client.Invoke("POST", "/aws/ha/volumes",
    map[string]string{
      "createAggregateIfNotFound": fmt.Sprint(createAggregateIfNotFound),
    },
    request,
  )
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}

// ModifyVolume modifies the given volume
func (api *AWSHAWorkingEnvironmentAPI) ModifyVolume(workenvId, svmName, volumeName string, request *workenv.VolumeModifyRequest) error {
  if request == nil {
		return errors.New(workenv.ErrInvalidVolumeCreationRequest)
	}

  _, err := api.Client.Invoke("PUT",
    fmt.Sprintf("/aws/ha/volumes/%v/%v/%v", workenvId, svmName, volumeName),
    nil, request)
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}

// DeleteVolume deletes the given volume
func (api *AWSHAWorkingEnvironmentAPI) DeleteVolume(workenvId, svmName, volumeName string) error {
  _, err := api.Client.Invoke("DELETE",
    fmt.Sprintf("/aws/ha/volumes/%v/%v/%v", workenvId, svmName, volumeName),
    nil, nil)
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}

// MoveVolume moves the given volume
func (api *AWSHAWorkingEnvironmentAPI) MoveVolume(workenvId, svmName, volumeName string, request *workenv.VolumeMoveRequest) error {
  if request == nil {
		return errors.New(workenv.ErrInvalidVolumeCreationRequest)
	}

  _, err := api.Client.Invoke("POST",
    fmt.Sprintf("/aws/ha/volumes/%v/%v/%v/move", workenvId, svmName, volumeName),
    nil, request)
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}

// CloneVolume clones the given volume
func (api *AWSHAWorkingEnvironmentAPI) CloneVolume(workenvId, svmName, volumeName string, request *workenv.VolumeCloneRequest) error {
  if request == nil {
		return errors.New(workenv.ErrInvalidVolumeCreationRequest)
	}

  _, err := api.Client.Invoke("POST",
    fmt.Sprintf("/aws/ha/volumes/%v/%v/%v/clone", workenvId, svmName, volumeName),
    nil, request)
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}

// ChangeVolumeDiskType changes disk type for the given volume
func (api *AWSHAWorkingEnvironmentAPI) ChangeVolumeDiskType(workenvId, svmName, volumeName string, request *workenv.VolumeChangeDiskTypeRequest) error {
  if request == nil {
		return errors.New(workenv.ErrInvalidVolumeCreationRequest)
	}

  _, err := api.Client.Invoke("POST",
    fmt.Sprintf("/aws/ha/volumes/%v/%v/%v/change-disk-type", workenvId, svmName, volumeName),
    nil, request)
  if err != nil {
		return errors.Wrap(err, client.ErrInvalidRequest)
	}

  return nil
}
