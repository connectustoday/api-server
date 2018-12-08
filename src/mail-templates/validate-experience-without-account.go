package mail_templates

var VALIDATE_EXPERIENCE_WITHOUT_ACCOUNT string

func init() {
	VALIDATE_EXPERIENCE_WITHOUT_ACCOUNT = `
<!DOCTYPE html>
<html>
<head>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
</head>
<body>
<div class="box gradient-45deg-indigo-blue gradient-shadow">
    <div style="text-align: center">Hello,
        <br>
        <br>
        A user on <a style="text-decoration:none; color: white" href="{{.Website}}">ConnectUS</a> has requested for a work/volunteer experience validation!</div>
    <br>
    These are the details of the request:
    <br>
    <br>
    <strong>Email:</strong> {{.Email}}
    <br>
    <strong>Username:</strong> {{.UserName}}
    <br>
    <strong>Full Name:</strong> {{.FullName}}
    <br>
    <strong>Name of Experience:</strong> {{.ExpName}}
    <br>
    <strong>Hours:</strong> {{.ExpHours}}
    <br>
    <strong>When it occurred:</strong> {{.ExpStart}} - {{.ExpEnd}}
    <br>
    <strong>Description:</strong>
    <br>
    {{expDesc}}
    <br>
    <hr>
    If you would like to approve this validation request, you can click the button below. Otherwise, you can safely disregard this email. 
    <br><br><br>
    <center>
        <a class="button gradient-45deg-light-blue-teal gradient-shadow" style="color:white; text-decoration:none" href='{{.VerifyLink}}'> <strong style="font-family: 'Nunito', sans-serif;">Approve</strong></a>
    </center>
    <br><br><br>
    <div style="text-align: center">Thank you!</div>
    <br>
</div>
Couldn't render the email? Click <a href="{{.VerifyLink}}">here.</a>
</body>
</html>
`
}