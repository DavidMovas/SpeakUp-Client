export namespace models {
	
	export class User {
	    id: string;
	    email: string;
	    username: string;
	    fullName: string;
	    avatarUrl: string;
	    bio: string;
	    createdAt: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.email = source["email"];
	        this.username = source["username"];
	        this.fullName = source["fullName"];
	        this.avatarUrl = source["avatarUrl"];
	        this.bio = source["bio"];
	        this.createdAt = source["createdAt"];
	    }
	}
	export class RegisterUserResponse {
	    accessToken: string;
	    user?: User;
	
	    static createFrom(source: any = {}) {
	        return new RegisterUserResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accessToken = source["accessToken"];
	        this.user = this.convertValues(source["user"], User);
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

