## List of endpoints
* /upload -> {POST} To upload any selected document
* /uploads -> {GET} To get a list of uploaded documents
* /download/:filename -> {GET} To download an uploaded file

### How to upload file
* Using Postman, select the POST http method, use ``localhost:8080/upload`` as the url
* For the body of the request, select ``form-data``
* A dropdown will show with title key, value. In the key use ``file`` and select File type, then select the file to be uploaded.

### Getting the list of uploaded file
* Change the HTTP method to GET and using the URL ``localhost:8080/uploaded``, a list of all uploaded file will be returned

### How to download any uploaded file
* Still using HTTP method GET and using the URL ``localhost:8080/download/:filename`` the :filename should be replaced with the name of the file to be downloaded.

### How to Run all test
Simply run the command
***go test -v ./...***