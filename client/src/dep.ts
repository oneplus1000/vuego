import { NetSvc } from "@/svc/net_svc";

export class Dep{
    static baseUrl :string = 'http://localhost:3000';
    static getNetSvc() : NetSvc {
        let net = new NetSvc(this.baseUrl);
        return net;
    }
} 