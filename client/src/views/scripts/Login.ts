import { Component, Vue } from 'vue-property-decorator';
import { Dep } from '@/dep';
import { LoginRespModel } from '@/model/login_resp_model';


@Component
export default class Login extends Vue {
    private username: string = '';
    private password: string = '';
    private msg: string = '';

    private async onLogin() {
        const net = Dep.getNetSvc();
        const json = await net.postJson(['api', 'login'], null);
        debugger;
    }

    private async resp(response: any): Promise<LoginRespModel> {
        const js = await response.text();
        const result = JSON.parse(js);
        const model = new LoginRespModel();
        model.userid = result.userid;
        model.result = result.result;
        return new Promise<LoginRespModel>((resolve) => {
            resolve(model);
        });
    }
}




