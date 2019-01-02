package core

//const .
const (
	INDEX = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
	<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
	<script src="https://use.fontawesome.com/4ef4b44616.js"></script>
</head>
<body>
<nav class="navbar navbar-light bg-primary mx-auto navbar-nav">
    <a class="navbar-brand" href="/" title="Do-Conv">
		<img src="assets/sunlight.png" style="width:40px"><span class="text-white"> Sunlight</span>
    </a>
</nav>
</br>
<div class="container" style="width: 700px">
  <div class="row">
    <div class="col-sm">
      <div class="card text-center" id="duration">
        <div class="card-body">
          <h3 class="card-title" id="">{{.Users}}</h3>
        </div>
        <div class="card-header">
			<i class="fa fa-users"></i> Users
		</div>
      </div>
    </div>
    <div class="col-sm">
      <div class="card text-center" id="target">
        <div class="card-body">
          <h3 class="card-title" id="">{{.NbrConv}} documents</h3>
        </div>
        <div class="card-header">
			<i class="fa fa-file"></i> Conversion
		</div>
      </div>
    </div>
  </div>
  <br>
    <form enctype="multipart/form-data" method="post" action="upload">
	  <h4>Email:</h4>
      <div class="input-group mb-2">
        <div class="input-group-prepend">
          <div class="input-group-text">@</div>
        </div>
        <input type="email" class="form-control" id="inlineFormInputGroup" placeholder="E-Mail" name="email" required>
      </div>
		<h4>Select the format of you'r conversion:</h4>
        <select class="custom-select custom-select-sm mb-3" name="type">
            <option value="docxtopdf">DOCX to PDF</option>
            <option value="pptxtopdf">PPTX to PDF</option>
            <option value="pdftojpg">PDF to Image</option>
            <option value="odttopdf">ODT to PDF</option>
            <option value="xlsxtopdf">XLSX to PDF</option>
            <option value="epubtopdf">EPUB to PDF</option>
        </select>
        </br>
		<h4>Choose a file:</h4>
        <div class="custom-file">
            <input type="file" class="custom-file-input" id="validatedCustomFile" required name="file">
            <label class="custom-file-label" for="validatedCustomFile" >Choose file...</label>
            <div class="invalid-feedback">Example invalid custom file feedback</div>
        </div>
		<input type="hidden" name="token" value="{{.Token}}"/>
        <br>
        <br>
        <button type="submit" class="btn btn-primary float-right">Convert <i class="fa fa-check"></i></button>
    </form>
	{{if .InfBool}}
		<div class="alert alert-warning alert-success fade show" role="alert">
			<strong>Info:</strong> {{.Inf}}
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
	{{else}}
	{{end}}
	{{if .ErrorBool}}
		<div class="alert alert-warning alert-dismissible fade show" role="alert">
			<strong>Error:</strong> {{.Error}}
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
	{{else}}
	{{end}}
</div>
<nav class="navbar navbar-dark bg-dark mt-5 fixed-bottom">
  <div class="navbar-expand m-auto navbar-text">
    Made with <i class="fa fa-heart text-danger"></i> by <a href="https://github.com/devectron">Devectron teams.</a>
  </div>
</nav>
</body>
</html>`
)
