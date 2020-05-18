import { JSONBytes } from '@tendermint/types';
import { AminoBytes } from './types';

/**
 * Encode a `PosmintTx` object from JSON to Amino
 *
 * @param   json           - JSON-encoded `StdTx` object
 * @param   lengthPrefixed - if `true`, use length-prefixed Amino encoding; if `false`, use bare Amino encoding
 *
 * @returns Amino-encoded `PosmintTx` object
 * @throws  will throw if encoding fails
 */
export function encodePosmintTx(json: JSONBytes, lengthPrefixed: boolean): AminoBytes;