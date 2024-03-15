# GO-Register

 ---
Go-Register is a full-stack, single-page application that implements user registration and login functions.
It has a clean, responsive UI, and great extensibility.

Originally I was making users register and log in from scratch for a website.
To make these functions easier to test and optimize, I decided to make a project exclusively for this. you can use this project for learning and demos.
 
 ---
## Before you start...
 1. Data storage:
    
    User Data are stored in a Hashmap, thus this project can run without the database.
    But keep in mind you should use a database or storage to persist data for your website or app. 
 
 3. JWT generate:
    
    The key in the auth.txt file is used to produce the JWT token, if you would like to use this project elsewhere
    , please make sure to change and hide the key to prevent security vulnerability.
    
 ---
## How to start
 
 1. First, clone this project:
    
    git clone https://github.com/guava-coder/go-register.git

 2. build the project in the terminal:
    
    ./build.sh

 3. go into the build folder and Run app.exe in the terminal:

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
 Email verification is necessary for functions like completing registration, changing email, etc...

 To enable the email verification function, you need to add a provider.json file, here is the example for Gmail:

    {
        "Sender": "yourgmail@gmail.com",
        "Token": "your password token",
        "Host": "smtp.gmail.com"
    }

 You can also use [MailTrap](https://mailtrap.io/), here is a [tutorial](https://mailtrap.io/blog/golang-send-email/).
 
 ---

|Functions|Currently Available|Set up Required|
|:-|:-|:-|
|Register|V|X (only store unauthorized user)|
|Email Verification|V|V|
|Login|V|X (default users only)|
|User Setting|V|X|
|Update Password|V|X|
 
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
