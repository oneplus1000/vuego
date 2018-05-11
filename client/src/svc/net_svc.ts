export class NetSvc {
    constructor(public baseUrl: string) { }
    private realUrl(paths: string[]): string {
        if (paths == null) {
            return this.baseUrl;
        }
        let url = this.baseUrl;
        let i = 0;
        const max = paths.length;
        while (i < max) {
            url += '/' + paths[i];
            i++;
        }
        return url;
    }

    public postJson(paths: string[], data: any): Promise<any> {
        return new Promise<any>((resolve, reject) => {
            let url = this.realUrl(paths);
            fetch(url,
                {
                    credentials: 'include',
                    //credentials: "same-origin",
                    method: "POST",
                    body: JSON.stringify(data)
                },
            ).then(async (resp) => {
                let json = await resp.json();
                resolve(json);
            }).catch((reason: any) => {
                reject(reason);
            });
        });
    }
}
