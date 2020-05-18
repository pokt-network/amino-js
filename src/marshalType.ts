import { jsonToBytes } from '@tendermint/belt';
import { AminoBytes } from '../lib/types';
import * as encodeType from './encodeType';
import { PosmintTx } from './types/pocket';

/**
 * Marshal a `Posmint` object to Amino
 *
 * @param   o              - `PosmintTx` object
 * @param   lengthPrefixed - if `true`, use length-prefixed Amino encoding; if `false`, use bare Amino encoding
 *
 * @returns Amino-encoded `PosmintTx` object
 * @throws  will throw if encoding fails
 */
export function marshalPosmintTx(o: PosmintTx, lengthPrefixed: boolean = true): AminoBytes {
    const json = jsonToBytes(o);
    return encodeType.encodePosmintTx(json, lengthPrefixed);
}