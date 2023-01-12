import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddBet } from "./types/lottery/lottery/tx";
import { MsgRevealBet } from "./types/lottery/lottery/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/lottery.lottery.MsgAddBet", MsgAddBet],
    ["/lottery.lottery.MsgRevealBet", MsgRevealBet],
    
];

export { msgTypes }