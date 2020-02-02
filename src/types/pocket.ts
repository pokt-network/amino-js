interface PosmintKey {
    /** Amino registered name, e.g. `"tendermint/PubKeySecp256k1"` */
    type: string;
    /** Base64-encoded key bytes */
    value: string;
}

/** @TODO document */
export interface PosmintPublicKey extends PosmintKey {
}

export interface PosmintEd25519PublicKey {
    
}

// /** @TODO document */
// export interface PosmintPubKey extends PosmintKey {
// }

/**
 * type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      uint64            `json:"sequence"`
}
 */

export interface PosmintEd25519PublicKey {
    /** Amino registered name, e.g. `"tendermint/PubKeySecp256k1"` */
    type: string;
    /** Base64-encoded key bytes */
    value: string;
}

export interface PosmintStdSignDoc {
    
    // account_number: number,
    // chain_id: string,
    // fee: string,
    // memo: string,
    // msgs: string[],
    // sequence: number
}

// export interface PosmintPubKey extends PosmintKey {
// }

export interface PosmintStdSignature {
    pub_key: PosmintEd25519PublicKey;
    signature: string;
}

export interface PosmintMsgSend {

}

export interface PosmintPublicKey {

}

export interface PosmintMsgStake {

}

export interface MsgBeginUnstake {

}

export interface MsgAppStake {

}

export interface MsgBeginAppUnstake {

}

/** @TODO document */
export interface PosmintMsg {
    type: string;
    value: PosmintMsgSend | PosmintMsgStake | MsgBeginUnstake | MsgAppStake | MsgBeginAppUnstake;
}

/** @TODO document */
export interface PosmintTx {
    type: string;
    value: PosmintTxValue;
}

/** @TODO document */
export interface PosmintTxValue {
}

/** @TODO document */
export interface PosmintCoin {
    denom: string;
    amount: string;
}

export interface PosmintStdTx extends PosmintTxValue {
    msg: PosmintMsg[];
    fee: PosmintCoin[];
    signatures: PosmintStdSignature[];
    memo: string;
}
