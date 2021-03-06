package mail_templates

const FORGOT_PASSWORD = `
<!DOCTYPE html>
<html>
<head>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <style>
        /* latin */
        @font-face {
            font-family: 'Nunito';
            font-style: normal;
            font-weight: 400;
            src: local('Nunito Regular'), local('Nunito-Regular'), url(https://fonts.gstatic.com/s/nunito/v9/XRXV3I6Li01BKofINeaB.woff2) format('woff2');
            unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
        }

        .box {
            font-family: 'Nunito', sans-serif;
            font-size: 17px;
            border-radius: 25px;
            width: 70%;
            margin: auto;
            padding: 30px 40px 30px 40px;
        }
        .button {
            padding: 10px 20px 10px 20px;
            text-decoration: none;
            color: white;
            box-shadow: 0 6px 20px 0 rgba(255, 248, 225, 0.5);
            border-radius: 5px;
            display: inline-block;
            width: 140px;
            border: none;
        }
        .gradient-45deg-light-blue-teal {
            background: #e91e63;
            background: -webkit-linear-gradient(45deg, #e91e63 0%, #f06292 100%);
            background: linear-gradient(45deg, #e91e63 0%, #f06292 100%);
        }
        .gradient-45deg-light-blue-teal.gradient-shadow {
            box-shadow: 0 6px 20px 0 rgba(100, 255, 218, 0.5);
        }
    </style>
</head>
<body>
<div class="box">
    <h3>Password Reset @ ConnectUS</h3>
    <hr>
    <br>
    Someone has requested a password reset on the account {{.Account}} at <a href="{{.Site}}">ConnectUS</a>.
    <br>
    <br>
    If this wasn't you, you can safely disregard this email.
    <br>
    <br>
    Otherwise, click the link below to continue on with the password reset process.
    <br>
    <br>
    <a class="button gradient-45deg-light-blue-teal gradient-shadow" style="color:white; text-decoration:none" href='{{.ResetLink}}'><strong style="font-family: 'Nunito', sans-serif;">Reset Password</strong></a>
    <br>
    <br>
    <hr>
</div>
</body>
</html>
`