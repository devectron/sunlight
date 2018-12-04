package core

const (
	INDEX = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>title</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
</head>
<body>
<nav class="navbar navbar-light bg-primary mx-auto navbar-nav">
    <a class="navbar-brand" href="/" title="Do-Conv">
        <span class="text-white">Do-Conv</span>
    </a>
</nav>
</br>
<div class="container" style="width: 800px">
    <form enctype="multipart/form-data" method="post" action="doconv">
        <select class="custom-select custom-select-sm mb-3" name="choice">
            <option value="0">PDF to DOC</option>
            <option value="1">DOC to PDF</option>
        </select>
        </br>
        <div class="custom-file">
            <input type="file" class="custom-file-input" id="validatedCustomFile" required name="file">
            <label class="custom-file-label" for="validatedCustomFile" >Choose file...</label>
            <div class="invalid-feedback">Example invalid custom file feedback</div>
        </div>
        </br>
        <button type="submit" class="btn btn-primary float-right">Send</button>
    </form>
    <p class=""></p>
    <p class=""></p>
</div>
</body>
</html>
`
)
