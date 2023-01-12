import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRevealBet } from "./types/lottery/lottery/tx";
import { MsgAddBet } from "./types/lottery/lottery/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/lottery.lottery.MsgRevealBet", MsgRevealBet],
    ["/lottery.lottery.MsgAddBet", MsgAddBet],
    
];

export { msgTypes }