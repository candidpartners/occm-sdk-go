// Package implements OCCM Working Environments API (VSA)
package vsa

import (
  "github.com/candidpartners/occm-sdk-go/api/workenv"
)

type VSAWorkingEnvironmentAPIProto interface {
	GetAggregates(string) ([]workenv.AggregateResponse, error)
  GetVolumes(string) ([]workenv.VolumeResponse, error)
  QuoteVolume(*VSAVolumeQuoteRequest) (*VSAVolumeQuoteResponse, error)
  CreateVolume(bool, *VSAVolumeCreateRequest) (error)
  ModifyVolume(string, string, string, *workenv.VolumeModifyRequest) error
  DeleteVolume(string, string, string) error
  MoveVolume(string, string, string, *workenv.VolumeMoveRequest) error
  CloneVolume(string, string, string, *workenv.VolumeCloneRequest) error
  ChangeVolumeDiskType(string, string, string, *workenv.VolumeChangeDiskTypeRequest) error
}

var _ VSAWorkingEnvironmentAPIProto = (*VSAWorkingEnvironmentAPI)(nil)
