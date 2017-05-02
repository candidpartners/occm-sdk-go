// Package implements OCCM Working Environments API (AWS-HA)
package awsha

import (
  "github.com/candidpartners/occm-sdk-go/api/workenv"
  "github.com/candidpartners/occm-sdk-go/api/workenv/vsa"
)

type AWSHAWorkingEnvironmentAPIProto interface {
  GetAggregates(string) ([]workenv.AggregateResponse, error)
  GetVolumes(string) ([]workenv.VolumeResponse, error)
  QuoteVolume(*vsa.VSAVolumeQuoteRequest) (*vsa.VSAVolumeQuoteResponse, error)
  CreateVolume(bool, *vsa.VSAVolumeCreateRequest) (error)
  ModifyVolume(string, string, string, *workenv.VolumeModifyRequest) error
  DeleteVolume(string, string, string) error
  MoveVolume(string, string, string, *workenv.VolumeMoveRequest) error
  CloneVolume(string, string, string, *workenv.VolumeCloneRequest) error
  ChangeVolumeDiskType(string, string, string, *workenv.VolumeChangeDiskTypeRequest) error
}

var _ AWSHAWorkingEnvironmentAPIProto = (*AWSHAWorkingEnvironmentAPI)(nil)
