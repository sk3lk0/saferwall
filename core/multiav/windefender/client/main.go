// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package windefender

import (
	"context"

	log "github.com/sirupsen/logrus"

	pb "github.com/saferwall/saferwall/core/multiav/windefender/proto"
	"google.golang.org/grpc"
)

const (
	address = "windefender-svc:50051"
)

// MultiAVScanResult av result
type MultiAVScanResult struct {
	Output   string `json:"output"`
	Infected bool   `json:"infected"`
}

// GetVerion returns version
func GetVerion(client pb.WinDefenderScannerClient) (*pb.VersionResponse, error) {
	versionRequest := &pb.VersionRequest{}
	return client.GetVersion(context.Background(), versionRequest)
}

// ScanFile scans file
func ScanFile(client pb.WinDefenderScannerClient, path string) (MultiAVScanResult, error) {
	log.Info("Scanning:", path)
	scanFile := &pb.ScanFileRequest{Filepath: path}
	res, err := client.ScanFile(context.Background(), scanFile)
	if err != nil {
		return MultiAVScanResult{}, err
	}

	return MultiAVScanResult{
		Output:   res.Output,
		Infected: res.Infected,
	}, nil
}

// Init connection
func Init() (pb.WinDefenderScannerClient, error) {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	conn, err := grpc.Dial(address, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}
	return pb.NewWinDefenderScannerClient(conn), nil
}
