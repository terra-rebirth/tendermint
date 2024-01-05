package config

import (
	"strings"
	"text/template"
)

// DefaultDirPerm is the default permissions used when creating directories.
//const DefaultDirPerm = 0o700

var genesisTemplate *template.Template
var gentexTemplate *template.Template
var privv_keyTemplate *template.Template

func init() {
	var err error
	//1
	tmpl := template.New("genesisFileTemplate").Funcs(template.FuncMap{
		"StringsJoin": strings.Join,
	})
	if genesisTemplate, err = tmpl.Parse(defaultGenesisTemplate); err != nil {
		panic(err)
	}
	//2
	tmpl = template.New("gentexFileTemplate").Funcs(template.FuncMap{
		"StringsJoin": strings.Join,
	})
	if gentexTemplate, err = tmpl.Parse(defaultGentexTemplate); err != nil {
		panic(err)
	}
	//3
	tmpl = template.New("privVKeyFileTemplate").Funcs(template.FuncMap{
		"StringsJoin": strings.Join,
	})
	if gentexTemplate, err = tmpl.Parse(defaultPrivValidatorKey); err != nil {
		panic(err)
	}
}

var defaultGentexTemplate = `{"body":{"messages":[{"@type":"/cosmos.staking.v1beta1.MsgCreateValidator","description":{"moniker":"core","identity":"","website":"","security_contact":"","details":""},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"min_self_delegation":"1000","delegator_address":"terra1gdd975kj5e49353f4emxwykjcqhs064tkt0q89","validator_address":"terravaloper1gdd975kj5e49353f4emxwykjcqhs064tkyrahk","pubkey":{"@type":"/cosmos.crypto.ed25519.PubKey","key":"tU2Comcl2A4+WLfQnEPsn+vd3K+XRqsvreZ/TdiR//s="},"value":{"denom":"uluna","amount":"100000000"}}],"memo":"62d1df73dc8554f81255c8aace50e691c09e632f@192.168.1.105:26656","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A/1tiSEKRuTJrrH86NoXC8+z0PNgBFeaFPD0yNGslyza"},"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"sequence":"0"}],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""},"tip":null},"signatures":["eiS7mVCYNCQjFWzNf39Qnm6NZB1w3vGl9Fs0fawd8XFxatfUkrGLz8GLTBvfOvTCN5srMFy+v1vt5bXrZC3TfA=="]}`

var defaultPrivValidatorKey = `{
	"address": "A8B912F0A73D64BFCDC4097F6E3FD9C9C102CDA9",
	"pub_key": {
	  "type": "tendermint/PubKeyEd25519",
	  "value": "tU2Comcl2A4+WLfQnEPsn+vd3K+XRqsvreZ/TdiR//s="
	},
	"priv_key": {
	  "type": "tendermint/PrivKeyEd25519",
	  "value": "l6SQxrgCaHQp0KqpN0uP8cqQmWTdKL+O12WywPZ1QFu1TYKiZyXYDj5Yt9CcQ+yf693cr5dGqy+t5n9N2JH/+w=="
	}
  }`

var defaultGenesisTemplate = `{
	"genesis_time": "2023-12-20T08:48:56.582756369Z",
	"chain_id": "paradise-1",
	"initial_height": "1",
	"consensus_params": {
	  "block": {
		"max_bytes": "22020096",
		"max_gas": "-1",
		"time_iota_ms": "1000"
	  },
	  "evidence": {
		"max_age_num_blocks": "100000",
		"max_age_duration": "172800000000000",
		"max_bytes": "1048576"
	  },
	  "validator": {
		"pub_key_types": [
		  "ed25519"
		]
	  },
	  "version": {}
	},
	"app_hash": "",
	"app_state": {
	  "alliance": {
		"params": {
		  "reward_delay_time": "604800s",
		  "take_rate_claim_interval": "300s",
		  "last_take_rate_claim_time": "0001-01-01T00:00:00Z"
		},
		"assets": [],
		"validator_infos": [],
		"reward_weight_change_snaphots": [],
		"delegations": [],
		"redelegations": [],
		"undelegations": []
	  },
	  "auth": {
		"params": {
		  "max_memo_characters": "256",
		  "tx_sig_limit": "7",
		  "tx_size_cost_per_byte": "10",
		  "sig_verify_cost_ed25519": "590",
		  "sig_verify_cost_secp256k1": "1000"
		},
		"accounts": [
		  {
			"@type": "/cosmos.vesting.v1beta1.ContinuousVestingAccount",
			"base_vesting_account": {
			  "base_account": {
				"address": "terra1gdd975kj5e49353f4emxwykjcqhs064tkt0q89",
				"pub_key": null,
				"account_number": "0",
				"sequence": "0"
			  },
			  "original_vesting": [
				{
				  "denom": "uluna",
				  "amount": "500000000000000"
				},
				{
				  "denom": "uusd",
				  "amount": "500000000000000"
				}
			  ],
			  "delegated_free": [],
			  "delegated_vesting": [],
			  "end_time": "1739667149"
			},
			"start_time": "1703667149"
		  }
		]
	  },
	  "authz": {
		"authorization": []
	  },
	  "bank": {
		"params": {
		  "send_enabled": [],
		  "default_send_enabled": true
		},
		"balances": [
		  {
			"address": "terra1gdd975kj5e49353f4emxwykjcqhs064tkt0q89",
			"coins": [
			  {
				"denom": "uluna",
				"amount": "1000000000000000"
			  },
			  {
				"denom": "uusd",
				"amount": "1000000000000000"
			  }
			]
		  }
		],
		"supply": [],
		"denom_metadata": []
	  },
	  "capability": {
		"index": "1",
		"owners": []
	  },
	  "crisis": {
		"constant_fee": {
		  "denom": "uluna",
		  "amount": "1000"
		}
	  },
	  "distribution": {
		"params": {
		  "community_tax": "0.020000000000000000",
		  "base_proposer_reward": "0.010000000000000000",
		  "bonus_proposer_reward": "0.040000000000000000",
		  "withdraw_addr_enabled": true
		},
		"fee_pool": {
		  "community_pool": []
		},
		"delegator_withdraw_infos": [],
		"previous_proposer": "",
		"outstanding_rewards": [],
		"validator_accumulated_commissions": [],
		"validator_historical_rewards": [],
		"validator_current_rewards": [],
		"delegator_starting_infos": [],
		"validator_slash_events": []
	  },
	  "evidence": {
		"evidence": []
	  },
	  "feegrant": {
		"allowances": []
	  },
	  "feeibc": {
		"identified_fees": [],
		"fee_enabled_channels": [],
		"registered_payees": [],
		"registered_counterparty_payees": [],
		"forward_relayers": []
	  },
	  "genutil": {
		"gen_txs": [
		  {
			"body": {
			  "messages": [
				{
				  "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
				  "description": {
					"moniker": "mainnet",
					"identity": "",
					"website": "",
					"security_contact": "",
					"details": ""
				  },
				  "commission": {
					"rate": "0.100000000000000000",
					"max_rate": "0.200000000000000000",
					"max_change_rate": "0.010000000000000000"
				  },
				  "min_self_delegation": "1000",
				  "delegator_address": "terra1gdd975kj5e49353f4emxwykjcqhs064tkt0q89",
				  "validator_address": "terravaloper1gdd975kj5e49353f4emxwykjcqhs064tkyrahk",
				  "pubkey": {
					"@type": "/cosmos.crypto.ed25519.PubKey",
					"key": "tU2Comcl2A4+WLfQnEPsn+vd3K+XRqsvreZ/TdiR//s="
				  },
				  "value": {
					"denom": "uluna",
					"amount": "100000000"
				  }
				}
			  ],
			  "memo": "62d1df73dc8554f81255c8aace50e691c09e632f@192.168.1.105:26656",
			  "timeout_height": "0",
			  "extension_options": [],
			  "non_critical_extension_options": []
			},
			"auth_info": {
			  "signer_infos": [
				{
				  "public_key": {
					"@type": "/cosmos.crypto.secp256k1.PubKey",
					"key": "A/1tiSEKRuTJrrH86NoXC8+z0PNgBFeaFPD0yNGslyza"
				  },
				  "mode_info": {
					"single": {
					  "mode": "SIGN_MODE_DIRECT"
					}
				  },
				  "sequence": "0"
				}
			  ],
			  "fee": {
				"amount": [],
				"gas_limit": "200000",
				"payer": "",
				"granter": ""
			  },
			  "tip": null
			},
			"signatures": [
			  "eiS7mVCYNCQjFWzNf39Qnm6NZB1w3vGl9Fs0fawd8XFxatfUkrGLz8GLTBvfOvTCN5srMFy+v1vt5bXrZC3TfA=="
			]
		  }
		]
	  },
	  "gov": {
		"starting_proposal_id": "1",
		"deposits": [],
		"votes": [],
		"proposals": [],
		"deposit_params": {
		  "min_deposit": [
			{
			  "denom": "uluna",
			  "amount": "10000000"
			}
		  ],
		  "max_deposit_period": "172800s"
		},
		"voting_params": {
		  "voting_period": "172800s"
		},
		"tally_params": {
		  "quorum": "0.334000000000000000",
		  "threshold": "0.500000000000000000",
		  "veto_threshold": "0.334000000000000000"
		}
	  },
	  "ibc": {
		"client_genesis": {
		  "clients": [],
		  "clients_consensus": [],
		  "clients_metadata": [],
		  "params": {
			"allowed_clients": [
			  "06-solomachine",
			  "07-tendermint"
			]
		  },
		  "create_localhost": false,
		  "next_client_sequence": "0"
		},
		"connection_genesis": {
		  "connections": [],
		  "client_connection_paths": [],
		  "next_connection_sequence": "0",
		  "params": {
			"max_expected_time_per_block": "30000000000"
		  }
		},
		"channel_genesis": {
		  "channels": [],
		  "acknowledgements": [],
		  "commitments": [],
		  "receipts": [],
		  "send_sequences": [],
		  "recv_sequences": [],
		  "ack_sequences": [],
		  "next_channel_sequence": "0"
		}
	  },
	  "ibchooks": {},
	  "interchainaccounts": {
		"controller_genesis_state": {
		  "active_channels": [],
		  "interchain_accounts": [],
		  "ports": [],
		  "params": {
			"controller_enabled": true
		  }
		},
		"host_genesis_state": {
		  "active_channels": [],
		  "interchain_accounts": [],
		  "port": "icahost",
		  "params": {
			"host_enabled": true,
			"allow_messages": [
			  "/cosmos.authz.v1beta1.MsgExec",
			  "/cosmos.authz.v1beta1.MsgGrant",
			  "/cosmos.authz.v1beta1.MsgRevoke",
			  "/cosmos.bank.v1beta1.MsgSend",
			  "/cosmos.bank.v1beta1.MsgMultiSend",
			  "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
			  "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
			  "/cosmos.distribution.v1beta1.MsgFundCommunityPool",
			  "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
			  "/cosmos.feegrant.v1beta1.MsgGrantAllowance",
			  "/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
			  "/cosmos.gov.v1beta1.MsgVoteWeighted",
			  "/cosmos.gov.v1beta1.MsgSubmitProposal",
			  "/cosmos.gov.v1beta1.MsgDeposit",
			  "/cosmos.gov.v1beta1.MsgVote",
			  "/cosmos.staking.v1beta1.MsgEditValidator",
			  "/cosmos.staking.v1beta1.MsgDelegate",
			  "/cosmos.staking.v1beta1.MsgUndelegate",
			  "/cosmos.staking.v1beta1.MsgBeginRedelegate",
			  "/cosmos.staking.v1beta1.MsgCreateValidator",
			  "/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
			  "/ibc.applications.transfer.v1.MsgTransfer",
			  "/cosmwasm.wasm.v1.MsgStoreCode",
			  "/cosmwasm.wasm.v1.MsgInstantiateContract",
			  "/cosmwasm.wasm.v1.MsgExecuteContract",
			  "/cosmwasm.wasm.v1.MsgMigrateContract"
			]
		  }
		}
	  },
	  "intertx": null,
	  "mint": {
		"minter": {
		  "inflation": "0.130000000000000000",
		  "annual_provisions": "0.000000000000000000"
		},
		"params": {
		  "mint_denom": "uluna",
		  "inflation_rate_change": "0.130000000000000000",
		  "inflation_max": "0.200000000000000000",
		  "inflation_min": "0.070000000000000000",
		  "goal_bonded": "0.670000000000000000",
		  "blocks_per_year": "6311520"
		}
	  },
	  "packetfowardmiddleware": {
		"params": {
		  "fee_percentage": "0.000000000000000000"
		},
		"in_flight_packets": {}
	  },
	  "params": null,
	  "slashing": {
		"params": {
		  "signed_blocks_window": "100",
		  "min_signed_per_window": "0.500000000000000000",
		  "downtime_jail_duration": "600s",
		  "slash_fraction_double_sign": "0.050000000000000000",
		  "slash_fraction_downtime": "0.010000000000000000"
		},
		"signing_infos": [],
		"missed_blocks": []
	  },
	  "staking": {
		"params": {
		  "unbonding_time": "1814400s",
		  "max_validators": 100,
		  "max_entries": 7,
		  "historical_entries": 10000,
		  "bond_denom": "uluna",
		  "min_commission_rate": "0.000000000000000000"
		},
		"last_total_power": "0",
		"last_validator_powers": [],
		"validators": [],
		"delegations": [],
		"unbonding_delegations": [],
		"redelegations": [],
		"exported": false
	  },
	  "tokenfactory": {
		"params": {
		  "denom_creation_fee": [
			{
			  "denom": "uluna",
			  "amount": "10000000"
			}
		  ]
		},
		"factory_denoms": []
	  },
	  "transfer": {
		"port_id": "transfer",
		"denom_traces": [],
		"params": {
		  "send_enabled": true,
		  "receive_enabled": true
		}
	  },
	  "upgrade": {},
	  "vesting": {},
	  "wasm": {
		"params": {
		  "code_upload_access": {
			"permission": "Everybody",
			"address": "",
			"addresses": []
		  },
		  "instantiate_default_permission": "Everybody"
		},
		"codes": [],
		"contracts": [],
		"sequences": [],
		"gen_msgs": []
	  }
	}
  }`
