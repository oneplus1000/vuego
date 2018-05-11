import { NetSvc } from "@/svc/net_svc";
import { Route } from "vue-router";
import router from "@/router";

export class GuardSvc {
    constructor(private netSvc: NetSvc) {

    }

    public beforeEnter(to: Route, from: Route, next: any) {
        //auth to server  
        this.netSvc.postJson(['api', 'is_logined'], {})
            .then((json: any) => {
                if (json.logined == true) {
                    next();
                } else {
                    router.push({ name: 'login' });
                }
            }).catch(() => {
                router.push({ name: 'login' });
            });
    }
}