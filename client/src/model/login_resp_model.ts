export class LoginRespModel {
    public result: string = '';
    public userid: string = '';

    public static fromJson(json :any) : LoginRespModel{
        const m = new LoginRespModel();
        m.result = json.result;
        m.userid = json.userid;
        return m;
    }
}