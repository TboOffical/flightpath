export namespace main {
	
	export class ServerInfo {
	    Name: string;
	    Addr: string;
	    Comments: string;
	    ApiKey: string;
	
	    static createFrom(source: any = {}) {
	        return new ServerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Addr = source["Addr"];
	        this.Comments = source["Comments"];
	        this.ApiKey = source["ApiKey"];
	    }
	}
	export class AppConfig {
	    Servers: ServerInfo[];
	    OnboardingStep: number;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Servers = this.convertValues(source["Servers"], ServerInfo);
	        this.OnboardingStep = source["OnboardingStep"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Param {
	    ParamName: string;
	    ParamType: string;
	    ParamDescription: string;
	    ParamJsonField: string;
	
	    static createFrom(source: any = {}) {
	        return new Param(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ParamName = source["ParamName"];
	        this.ParamType = source["ParamType"];
	        this.ParamDescription = source["ParamDescription"];
	        this.ParamJsonField = source["ParamJsonField"];
	    }
	}
	
	export class Task {
	    Name: string;
	    Params: Param[];
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Params = this.convertValues(source["Params"], Param);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace struct { ID int "json:\"ID\""; CreatedAt string "json:\"CreatedAt\""; UpdatedAt string "json:\"UpdatedAt\""; DeletedAt interface {} "json:\"DeletedAt\""; Data string "json:\"Data\"" } {
	
	export class  {
	    ID: number;
	    CreatedAt: string;
	    UpdatedAt: string;
	    DeletedAt: any;
	    Data: string;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = source["CreatedAt"];
	        this.UpdatedAt = source["UpdatedAt"];
	        this.DeletedAt = source["DeletedAt"];
	        this.Data = source["Data"];
	    }
	}

}

