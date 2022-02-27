
window.getImage = function () {
  window.go.main.App.GetImageInfo().then((res)=>{
    let bingImage = document.getElementById("bingImage");
    bingImage.setAttribute("src",res.url)
    bingImage.setAttribute("alt",res.title)
    bingImage.setAttribute("image-date",res.date)
    bingImage.hidden = false
  })
}

window.setWallpaper = function(){
  let bingImage = document.getElementById("bingImage");
  let url = bingImage.getAttribute("src");
  window.go.main.App.SetWallpaper(url).then(r => function(res){
    console.log(res)
  });
}

window.onload = function (){
  getImage()
}