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

export interface PosmintEd25519PublicKey {
    /** Amino registered name, e.g. `"tendermint/PubKeySecp256k1"` */
    type: string;
    /** Base64-encoded key bytes */
    value: string;
}

export interface PosmintStdSignDoc {
}

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

export interface PosmintMsgBeginUnstake {

}

export interface PosmintMsgUnjail {

}

export interface PosmintMsgAppStake {

}

export interface PosmintMsgBeginAppUnstake {

}

export interface PosmintMsgAppUnjail {

}

/** @TODO document */
export interface PosmintMsg {
    type: string;
    value: PosmintMsgSend | PosmintMsgStake | PosmintMsgBeginUnstake | PosmintMsgUnjail | PosmintMsgAppStake | PosmintMsgBeginAppUnstake | PosmintMsgAppUnjail;
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
    msg: PosmintMsg;
    fee: PosmintCoin[];
    signature: PosmintStdSignature;
    memo: string;
    entropy: string;
}
