import { NetSvc } from "@/svc/net_svc";
import { GuardSvc } from "@/svc/guard_svc";

export class Dep{
    static baseUrl :string = 'http://192.168.1.3:3000';
    static getNetSvc() : NetSvc {
        let net = new NetSvc(this.baseUrl);
        return net;
    }
    static getGuardSvc() : GuardSvc{
        let gd = new GuardSvc(this.getNetSvc());
        return gd; 
    }
} 