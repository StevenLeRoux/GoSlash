package main 

const (
alias = `<!doctype html>
<html lang="en-US">

<head>
	<meta charset="utf-8">
	<meta http-equiv="x-ua-compatible" content="ie=edge">
	<title>GoSlash</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="robots" content="noindex">
	<link href="https://fonts.googleapis.com/css?family=Bitter|Source+Sans+Pro" rel="stylesheet">
	<style>
body {
  background: #3a7bd5;
  background: -webkit-linear-gradient(to left, #3a6073, #3a7bd5);
  background: -webkit-linear-gradient(right, #3a6073, #3a7bd5);
  background: linear-gradient(to left, #3a6073, #3a7bd5);
}
body h1 {
  font-family: 'Bitter', serif;
}
body .inner {
  max-width: 1024px;
  margin: 0 auto;
}
body .inner #alias-list {
  background: #ffffff;
  border-radius: 3px;
  overflow: hidden;
}
body .inner #alias-list .alias:nth-child(1) {
  border-top: none;
}
body .inner #alias-list .alias {
  border-top: 1px solid #333;
  padding: 16px;
  position: relative;
  cursor: pointer;
}
body .inner #alias-list .alias h3 {
  display: inline-block;
  font-family: 'Bitter', serif;
  margin: 0;
  padding: 0 0 4px;
}
body .inner #alias-list .alias p {
  font-family: 'Source Sans Pro', sans-serif;
  margin: 0;
  padding: 0;
}
body .inner #alias-list .alias .target {
  display: inline-block;
}
body .inner #alias-list .alias .icon {
  position: absolute;
  top: 16px;
  right: 16px;
  bottom: 16px;
  height: 32px;
  margin: auto;
}
body .inner #alias-list .alias:hover {
  background: #EEE;
}

/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbImFsaWFzLmxlc3MiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUE7RUFDRSxvQkFBb0I7RUFDcEIsK0RBQStEO0VBQy9ELDZEQUF1RDtFQUF2RCx1REFBdUQ7Q0FDeEQ7QUFDRDtFQUNFLDZCQUE2QjtDQUM5QjtBQUNEO0VBQ0Usa0JBQWtCO0VBQ2xCLGVBQWU7Q0FDaEI7QUFDRDtFQUNFLG9CQUFvQjtFQUNwQixtQkFBbUI7RUFDbkIsaUJBQWlCO0NBQ2xCO0FBQ0Q7RUFDRSxpQkFBaUI7Q0FDbEI7QUFDRDtFQUNFLDJCQUEyQjtFQUMzQixjQUFjO0VBQ2QsbUJBQW1CO0VBQ25CLGdCQUFnQjtDQUNqQjtBQUNEO0VBQ0Usc0JBQXNCO0VBQ3RCLDZCQUE2QjtFQUM3QixVQUFVO0VBQ1YsaUJBQWlCO0NBQ2xCO0FBQ0Q7RUFDRSwyQ0FBMkM7RUFDM0MsVUFBVTtFQUNWLFdBQVc7Q0FDWjtBQUNEO0VBQ0Usc0JBQXNCO0NBQ3ZCO0FBQ0Q7RUFDRSxtQkFBbUI7RUFDbkIsVUFBVTtFQUNWLFlBQVk7RUFDWixhQUFhO0VBQ2IsYUFBYTtFQUNiLGFBQWE7Q0FDZDtBQUNEO0VBQ0UsaUJBQWlCO0NBQ2xCIiwiZmlsZSI6ImFsaWFzLmxlc3MiLCJzb3VyY2VzQ29udGVudCI6WyJib2R5IHtcbiAgYmFja2dyb3VuZDogIzNhN2JkNTtcbiAgYmFja2dyb3VuZDogLXdlYmtpdC1saW5lYXItZ3JhZGllbnQodG8gbGVmdCwgIzNhNjA3MywgIzNhN2JkNSk7XG4gIGJhY2tncm91bmQ6IGxpbmVhci1ncmFkaWVudCh0byBsZWZ0LCAjM2E2MDczLCAjM2E3YmQ1KTtcbn1cbmJvZHkgaDEge1xuICBmb250LWZhbWlseTogJ0JpdHRlcicsIHNlcmlmO1xufVxuYm9keSAuaW5uZXIge1xuICBtYXgtd2lkdGg6IDEwMjRweDtcbiAgbWFyZ2luOiAwIGF1dG87XG59XG5ib2R5IC5pbm5lciAjYWxpYXMtbGlzdCB7XG4gIGJhY2tncm91bmQ6ICNmZmZmZmY7XG4gIGJvcmRlci1yYWRpdXM6IDNweDtcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcbn1cbmJvZHkgLmlubmVyICNhbGlhcy1saXN0IC5hbGlhczpudGgtY2hpbGQoMSkge1xuICBib3JkZXItdG9wOiBub25lO1xufVxuYm9keSAuaW5uZXIgI2FsaWFzLWxpc3QgLmFsaWFzIHtcbiAgYm9yZGVyLXRvcDogMXB4IHNvbGlkICMzMzM7XG4gIHBhZGRpbmc6IDE2cHg7XG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcbiAgY3Vyc29yOiBwb2ludGVyO1xufVxuYm9keSAuaW5uZXIgI2FsaWFzLWxpc3QgLmFsaWFzIGgzIHtcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xuICBmb250LWZhbWlseTogJ0JpdHRlcicsIHNlcmlmO1xuICBtYXJnaW46IDA7XG4gIHBhZGRpbmc6IDAgMCA0cHg7XG59XG5ib2R5IC5pbm5lciAjYWxpYXMtbGlzdCAuYWxpYXMgcCB7XG4gIGZvbnQtZmFtaWx5OiAnU291cmNlIFNhbnMgUHJvJywgc2Fucy1zZXJpZjtcbiAgbWFyZ2luOiAwO1xuICBwYWRkaW5nOiAwO1xufVxuYm9keSAuaW5uZXIgI2FsaWFzLWxpc3QgLmFsaWFzIC50YXJnZXQge1xuICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XG59XG5ib2R5IC5pbm5lciAjYWxpYXMtbGlzdCAuYWxpYXMgLmljb24ge1xuICBwb3NpdGlvbjogYWJzb2x1dGU7XG4gIHRvcDogMTZweDtcbiAgcmlnaHQ6IDE2cHg7XG4gIGJvdHRvbTogMTZweDtcbiAgaGVpZ2h0OiAzMnB4O1xuICBtYXJnaW46IGF1dG87XG59XG5ib2R5IC5pbm5lciAjYWxpYXMtbGlzdCAuYWxpYXM6aG92ZXIge1xuICBiYWNrZ3JvdW5kOiAjRUVFO1xufVxuIl19 */	</style>

	<link href="https://fonts.googleapis.com/css?family=Bitter|Source+Sans+Pro:400,700" rel="stylesheet">
</head>

<body>
	<div class="inner">
		<h1>GoSlash</h1>
		<div id="alias-list"></div>
	</div>
	<script>
!function(){function a(){var a=document.getElementById("alias-list");b.forEach(function(b){console.log(b);var c=document.createElement("div");c.classList.add("alias"),c.innerHTML="<h3>"+b.alias+'</h3><p class="target">&nbsp;-&nbsp;'+b.target+"</p><p>"+b.description+'</p><span class="icon"><?xml version="1.0" encoding="UTF-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="32" height="32" viewBox="0 0 24 24"><path d="M4,11V13H16L10.5,18.5L11.92,19.92L19.84,12L11.92,4.08L10.5,5.5L16,11H4Z" /></svg></span>',c.addEventListener("click",function(){window.location.href="/"+b.alias}),a.appendChild(c)})}var b=[],c=new XMLHttpRequest;c.onload=function(c){b=c.target.response,a()},c.open("GET","http://127.0.0.1:8080/dump?"+(new Date).getTime(),!0),c.responseType="json",c.send()}();	</script>
</body>

</html>
`
)
