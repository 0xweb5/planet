import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSendUpdatePost } from "./types/planet/blog/tx";
import { MsgSendIbcPost } from "./types/planet/blog/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/planet.blog.MsgSendUpdatePost", MsgSendUpdatePost],
    ["/planet.blog.MsgSendIbcPost", MsgSendIbcPost],
    
];

export { msgTypes }