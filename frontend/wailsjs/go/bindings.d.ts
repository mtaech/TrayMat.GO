export interface go {
  "main": {
    "App": {
		GetImageInfo():Promise<ImageInfo>
		Greet(arg1:string):Promise<string>
		SetWallpaper(arg1:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
