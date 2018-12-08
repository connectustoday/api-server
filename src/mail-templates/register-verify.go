package mail_templates

var REGISTER_VERIFY string

func init() {
	REGISTER_VERIFY = `
<!DOCTYPE html>
<html>
<head>
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
            color: white;
            font-size: 18px;
            border-radius: 25px;
            text-align: center;
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
            width: 100px;
            border: none;
        }
        .gradient-45deg-light-blue-teal {
            background: #b3e5fc;
            background: -webkit-linear-gradient(45deg, #b3e5fc 0%, #09a2ce 100%);
            background: linear-gradient(45deg, #b3e5fc 0%, #09a2ce 100%);
        }
        .gradient-45deg-light-blue-teal.gradient-shadow {
            box-shadow: 0 6px 20px 0 rgba(100, 255, 218, 0.5);
        }
        .gradient-45deg-indigo-blue {
            background: #303f9f;
            background: -webkit-linear-gradient(45deg, #303f9f 0%, #1976D2 100%);
            background: linear-gradient(45deg, #303f9f 0%, #1976D2 100%);
        }
        .gradient-45deg-indigo-blue.gradient-shadow {
            box-shadow: 0 6px 20px 0 rgba(25, 118, 210, 0.5);
        }
    </style>
</head>
<body>
<div class="box gradient-45deg-indigo-blue gradient-shadow">
    <strong>Thank you for signing up for ConnectUS!</strong>
    <br>
    In order to activate your account, please click the button below:
    <br><br>
    <a class="button gradient-45deg-light-blue-teal gradient-shadow" style="color:white; text-decoration:none" href='{{.VerifyLink}}'><strong style="font-family: 'Nunito', sans-serif;">Activate</strong></a>
    <br><br>
</div>
</body>
</html>

`
}