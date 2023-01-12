import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRevealBet } from "./types/lottery/lottery/tx";
import { MsgUpdateLottery } from "./types/lottery/lottery/tx";
import { MsgAddBet } from "./types/lottery/lottery/tx";
import { MsgCreateLottery } from "./types/lottery/lottery/tx";
import { MsgDeleteLottery } from "./types/lottery/lottery/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/lottery.lottery.MsgRevealBet", MsgRevealBet],
    ["/lottery.lottery.MsgUpdateLottery", MsgUpdateLottery],
    ["/lottery.lottery.MsgAddBet", MsgAddBet],
    ["/lottery.lottery.MsgCreateLottery", MsgCreateLottery],
    ["/lottery.lottery.MsgDeleteLottery", MsgDeleteLottery],
    
];

export { msgTypes }