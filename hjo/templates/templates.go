package templates


const WELCOME_EMAIL = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .email-container {
            background-color: #ffffff;
            margin: 20px auto;
            padding: 20px;
            max-width: 600px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .header {
            background-color: #ff0037;
            color: white;
            padding: 10px;
            text-align: center;
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
        }
        .content {
            padding: 20px;
            text-align: center;
        }
        .footer {
            background-color: #f1f1f1;
            color: #555;
            padding: 10px;
            text-align: center;
            border-bottom-left-radius: 10px;
            border-bottom-right-radius: 10px;
        }
        a {
            color: #007bff;
            text-decoration: none;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="header">
            <h1>Heya {{ .FullName }}!</h1>
        </div>
        <div class="content">
            <p>Welcome to GoEas!</p>
            <p>Don't forget to star our project on GitHub:</p>
            <a href="{{ .Link }}">{{ .Link }}</a>
        </div>
        <div class="footer">
            <p>Thank you for joining us!</p>
        </div>
    </div>
</body>
</html>`
