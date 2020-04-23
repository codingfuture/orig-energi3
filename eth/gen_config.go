// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/consensus/ethash"
	"energi.world/core/gen3/core"
	"energi.world/core/gen3/eth/downloader"
	"energi.world/core/gen3/eth/gasprice"
)

var _ = (*configMarshaling)(nil)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		NoPruning               bool
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               int                    `toml:",omitempty"`
		LightPeers              int                    `toml:",omitempty"`
		SkipBcVersionCheck      bool                   `toml:"-"`
		DatabaseHandles         int                    `toml:"-"`
		DatabaseCache           int
		TrieCleanCache          int
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		TrieRapidTime           time.Duration
		Etherbase               common.Address `toml:",omitempty"`
		MinerNotify             []string       `toml:",omitempty"`
		MinerExtraData          hexutil.Bytes  `toml:",omitempty"`
		MinerGasFloor           uint64
		MinerGasCeil            uint64
		MinerGasPrice           *big.Int
		MinerRecommit           time.Duration
		MinerNoverify           bool
		MinerDPoS               DPoSMap `toml:",omitempty"`
		MinerMigration          string  `toml:",omitempty"`
		MinerNonceCap           uint64  `toml:"-"`
		MinerBailout            uint64  `toml:"-"`
		MinerAutocollateral     uint64  `toml:",omitempty"`
		PublicService           bool    `toml:",omitempty"`
		Ethash                  ethash.Config
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
		EWASMInterpreter        string
		EVMInterpreter          string
		ConstantinopleOverride  *big.Int
		RPCGasCap               *big.Int `toml:",omitempty"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.NoPruning = c.NoPruning
	enc.Whitelist = c.Whitelist
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.TrieRapidTime = c.TrieRapidTime
	enc.Etherbase = c.Etherbase
	enc.MinerNotify = c.MinerNotify
	enc.MinerExtraData = c.MinerExtraData
	enc.MinerGasFloor = c.MinerGasFloor
	enc.MinerGasCeil = c.MinerGasCeil
	enc.MinerGasPrice = c.MinerGasPrice
	enc.MinerRecommit = c.MinerRecommit
	enc.MinerNoverify = c.MinerNoverify
	enc.MinerDPoS = c.MinerDPoS
	enc.MinerMigration = c.MinerMigration
	enc.MinerNonceCap = c.MinerNonceCap
	enc.MinerBailout = c.MinerBailout
	enc.MinerAutocollateral = c.MinerAutocollateral
	enc.PublicService = c.PublicService
	enc.Ethash = c.Ethash
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.EWASMInterpreter = c.EWASMInterpreter
	enc.EVMInterpreter = c.EVMInterpreter
	enc.ConstantinopleOverride = c.ConstantinopleOverride
	enc.RPCGasCap = c.RPCGasCap
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		NoPruning               *bool
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               *int                   `toml:",omitempty"`
		LightPeers              *int                   `toml:",omitempty"`
		SkipBcVersionCheck      *bool                  `toml:"-"`
		DatabaseHandles         *int                   `toml:"-"`
		DatabaseCache           *int
		TrieCleanCache          *int
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		TrieRapidTime           *time.Duration
		Etherbase               *common.Address `toml:",omitempty"`
		MinerNotify             []string        `toml:",omitempty"`
		MinerExtraData          *hexutil.Bytes  `toml:",omitempty"`
		MinerGasFloor           *uint64
		MinerGasCeil            *uint64
		MinerGasPrice           *big.Int
		MinerRecommit           *time.Duration
		MinerNoverify           *bool
		MinerDPoS               *DPoSMap `toml:",omitempty"`
		MinerMigration          *string  `toml:",omitempty"`
		MinerNonceCap           *uint64  `toml:"-"`
		MinerBailout            *uint64  `toml:"-"`
		MinerAutocollateral     *uint64  `toml:",omitempty"`
		PublicService           *bool    `toml:",omitempty"`
		Ethash                  *ethash.Config
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
		EWASMInterpreter        *string
		EVMInterpreter          *string
		ConstantinopleOverride  *big.Int
		RPCGasCap               *big.Int `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.TrieRapidTime != nil {
		c.TrieRapidTime = *dec.TrieRapidTime
	}
	if dec.Etherbase != nil {
		c.Etherbase = *dec.Etherbase
	}
	if dec.MinerNotify != nil {
		c.MinerNotify = dec.MinerNotify
	}
	if dec.MinerExtraData != nil {
		c.MinerExtraData = *dec.MinerExtraData
	}
	if dec.MinerGasFloor != nil {
		c.MinerGasFloor = *dec.MinerGasFloor
	}
	if dec.MinerGasCeil != nil {
		c.MinerGasCeil = *dec.MinerGasCeil
	}
	if dec.MinerGasPrice != nil {
		c.MinerGasPrice = dec.MinerGasPrice
	}
	if dec.MinerRecommit != nil {
		c.MinerRecommit = *dec.MinerRecommit
	}
	if dec.MinerNoverify != nil {
		c.MinerNoverify = *dec.MinerNoverify
	}
	if dec.MinerDPoS != nil {
		c.MinerDPoS = *dec.MinerDPoS
	}
	if dec.MinerMigration != nil {
		c.MinerMigration = *dec.MinerMigration
	}
	if dec.MinerNonceCap != nil {
		c.MinerNonceCap = *dec.MinerNonceCap
	}
	if dec.MinerBailout != nil {
		c.MinerBailout = *dec.MinerBailout
	}
	if dec.MinerAutocollateral != nil {
		c.MinerAutocollateral = *dec.MinerAutocollateral
	}
	if dec.PublicService != nil {
		c.PublicService = *dec.PublicService
	}
	if dec.Ethash != nil {
		c.Ethash = *dec.Ethash
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.EWASMInterpreter != nil {
		c.EWASMInterpreter = *dec.EWASMInterpreter
	}
	if dec.EVMInterpreter != nil {
		c.EVMInterpreter = *dec.EVMInterpreter
	}
	if dec.ConstantinopleOverride != nil {
		c.ConstantinopleOverride = dec.ConstantinopleOverride
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = dec.RPCGasCap
	}
	return nil
}
