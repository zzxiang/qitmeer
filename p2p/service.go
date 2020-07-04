/*
 * Copyright (c) 2020.
 * Project:qitmeer
 * File:Service.go
 * Date:7/2/20 8:04 PM
 * Author:Jin
 * Email:lochjin@gmail.com
 */
package p2p

import (
	"context"
	"crypto/ecdsa"
	"github.com/Qitmeer/qitmeer/config"
	"github.com/Qitmeer/qitmeer/p2p/peers"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"time"
)

type Service struct {
	started               bool
	isPreGenesis          bool
	pingMethod            func(ctx context.Context, id peer.ID) error
	cancel                context.CancelFunc
	peers                 *peers.Status
	privKey               *ecdsa.PrivateKey
	startupErr            error
	ctx                   context.Context
	host                  host.Host
	genesisTime           time.Time
	genesisValidatorsRoot []byte
}

func (s *Service) Start() error {
	log.Info("P2P Service Start")
	return nil
}

func (s *Service) Stop() error {
	log.Info("P2P Service Stop")
	return nil
}

func NewService(cfg *config.Config) (*Service, error) {
	s := Service{}

	return &s, nil
}
