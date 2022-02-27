/* Do not change, this code is generated from Golang structs */

export {};

export class ImageInfo {
    url: string;
    date: string;
    title: string;

    static createFrom(source: any = {}) {
        return new ImageInfo(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.url = source["url"];
        this.date = source["date"];
        this.title = source["title"];
    }
}