export namespace models {
	
	export class Card {
	    Name: string;
	    ImagePath: string;
	    Count: number;
	    Set: string;
	    SetNumber: string;
	
	    static createFrom(source: any = {}) {
	        return new Card(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.ImagePath = source["ImagePath"];
	        this.Count = source["Count"];
	        this.Set = source["Set"];
	        this.SetNumber = source["SetNumber"];
	    }
	}

}

