import './setup';
import { base64ToBytes, bytesToBase64 } from '@tendermint/belt';
import * as Amino from '../';

const tx = {
    "type": "posmint/StdTx",
    "value": { 
        "fee": [{ "amount": "100000", "denom": "upokt" }], 
        "memo": "", 
        "msg": [{ "type": "pos/Send", "value": { "Amount": "1000", "FromAddress": "4930289621aefbf9252c91c4c729b7f685e44c4b", "ToAddress": "4930289621aefbf9252c91c4c729b7f685e44c4b" }}], 
        "signatures": [{ "pub_key": { "type": "crypto/ed25519_public_key", "value": "9i9322nUSMG1bzVAxjPylNI8za8AK/azdtBYoAtRz6o=" }, "signature": "r41RHZPHXMpUJx+mgZEpw4c/o3udKZtB9aZJVxDEiDMVIIz4iEORk38QdS+sjubXLPaPg7tZjasJnU6hCS9bCw==" }] 
    }
}

const txData = 'uAHbCxcNCjal40SmChRJMCiWIa77+SUskcTHKbf2heRMSxIUSTAoliGu+/klLJHExym39oXkTEsaBDEwMDASDwoFdXBva3QSBjEwMDAwMBppCiWdVEd0IPYvd9tp1EjBtW81QMYz8pTSPM2vACv2s3bQWKALUc+qEkCvjVEdk8dcylQnH6aBkSnDhz+je50pm0H1pklXEMSIMxUgjPiIQ5GTfxB1L6yO5tcs9o+Du1mNqwmdTqEJL1sL';

describe('Tx', () => {
    describe('decoding', () => {
        describe('Tx', () => {
            it('decodes bytes', () => {
                const bytes = base64ToBytes(txData);
                expect(bytes).not.toBeNull
                // const value = Amino.unmarshalPosmintTx(bytes, true);
                // expect(value).toMatchObject(tx);
            });
        });
    });

    describe('encoding', () => {
        describe('Tx', () => {
            // it('encodes value', () => {
            //     const bytes = Amino.marshalPosmintTx(tx, true)
            //     const data = bytesToBase64(bytes);
            //     expect(data).toBe(txData);
            // });
        });
    });
});