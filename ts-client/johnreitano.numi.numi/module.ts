// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgCreateAndVerifyUser } from "./types/numi/numi/tx";


export { MsgCreateAndVerifyUser };

type sendMsgCreateAndVerifyUserParams = {
  value: MsgCreateAndVerifyUser,
  fee?: StdFee,
  memo?: string
};


type msgCreateAndVerifyUserParams = {
  value: MsgCreateAndVerifyUser,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgCreateAndVerifyUser({ value, fee, memo }: sendMsgCreateAndVerifyUserParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAndVerifyUser: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAndVerifyUser({ value: MsgCreateAndVerifyUser.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAndVerifyUser: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgCreateAndVerifyUser({ value }: msgCreateAndVerifyUserParams): EncodeObject {
			try {
				return { typeUrl: "/johnreitano.numi.numi.MsgCreateAndVerifyUser", value: MsgCreateAndVerifyUser.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAndVerifyUser: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			JohnreitanoNumiNumi: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;