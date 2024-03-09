# GO-Register

 ---
Go-Register is a full-stack, single-page website that implements user registration and login features.
It has a clean, responsive UI, and great extensibility.

If you are...

 * A web dev who wants the user register feature.
 * A designer who wants to make UI/UX around user login.
 * Or doesn't want to make these features from scratch :P
 
 Then this is for you! 
 
 ---
## Before you start...
 1. Data storage:
    
    User Data are stored in a Hashmap, thus the app can start without the database. But keep in mind using a database to persist data is very important. 
 
 2. JWT generate:
    
    The key in the auth.txt file is used to produce the JWT token, if you want to use this project elsewhere, make sure to change and hide the key to prevent security vulnerability.
    
 ---
## How to start
 
 1. First, clone this project:
    
    git clone https://github.com/guava-coder/go-register.git

 2. build the project in ternimal:
    
    ./build.sh

 3. go into build folder and Run app.exe in terminal:

    cd build/go-register/

    ./app.exe

 4. After running, open the app in the browser with the URL below:

    http://localhost:8082/

## Login
 By default, there is an authorized user, using it to log in successfully:
 
    Email: mark@mail.com
    Password: 123

 There is also an unauthorized user, using it to try login failed:

    Email: lisa@mail.com
    Password: 123

## Set Up Email Verification
 Email verification is necessary for features like completing registration, changing email, etc...

 To enable the email verification feature, you need to add a provider.json file, here is the example for Gmail:

    {
        "Sender": "yourgmail@gmail.com",
        "Token": "your password token",
        "Host": "smtp.gmail.com"
    }

 You can also use [MailTrap](https://mailtrap.io/), here is a [tutorial](https://mailtrap.io/blog/golang-send-email/).
 
 ---

|Features|Currently Available|Set up Required|
|:-|:-|:-|
|Register|V|X (only store unauthorized user)|
|Email Verification|V|V|
|Login|V|X (default users only)|
|User Setting|V|X|
|Update Email|X|V|
 
 --- 
 ## Open Source Dependencies

 ### Backend

 * [email-verifier](https://github.com/AfterShip/email-verifier)

 * [Gin Web Framework](https://github.com/gin-gonic/gin)

 * [jwt-go](https://github.com/golang-jwt/jwt)

 * [google/uuid](https://github.com/google/uuid)

 ### Frontend

 * [HTMX](https://github.com/bigskysoftware/htmx)

 * [Bootstrap 5](https://github.com/twbs/bootstrap)

 * [JavaScript Cookie](https://github.com/js-cookie/js-cookie)
