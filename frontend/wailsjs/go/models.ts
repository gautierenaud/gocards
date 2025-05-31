export namespace models {
	
	export class Card {
	    Name: string;
	    ImagePath: string;
	
	    static createFrom(source: any = {}) {
	        return new Card(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.ImagePath = source["ImagePath"];
	    }
	}

}

