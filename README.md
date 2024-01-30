# GO-Register

### Register/Login web app, Restiful APIs and JWT 

 ---
 > Go-Register is a full-stack, single-page web app provides features around user register and login.

 > It has a clean, responsive UI, and great extensibility.

 > If you are...

 * A web dev who wants to learn the user register feature.
 * A designer who wants to make UI/UX around user login.
 * Or doesn't want to make these features from scratch :P
 
 > Then this is for you! 
 ---
 ## Before you start...
 1. Data storage
   >> User Data are stored in hashmap, thus the app can start without the database. But keep in mind using a database to persist data is very important. 
 2. JWT generate
   >> JWT are produced by the key in auth.txt file, if you want to use this project elsewhere, remember to change and hide the key to prevent security vulnerability. 
 ---
## How to start

 > First, clone this project:
    
    git clone https://github.com/guava-coder/go-register.git

 > Run app.exe in terminal:

    ./app.exe

 > After run, open the app in browser with url below:

    http://localhost:8082

## Login
 > By default, there is one authorized user, you can use it to log in successfully:
 
    Email:    "mark@mail.com",
	Password: "123"

 > There is also one unauthorized user, you can use it to test login failed:

    Email:    "lisa@mail.com",
    Password: "123"

## Set Up Email Verification
 > Email verification is necessary for some features like completing registration, changing email, etc...

 > To enable email verification feature, you need to add provider.json file, here is the example for gmail:

    {
        "Sender":"yourgmail@gmail.com",
        "Token":"your password token",
        "Host":"smtp.gmail.com"
    }

 > You can also use [MailTrap](https://mailtrap.io/), you can find tutorial [here](https://mailtrap.io/blog/golang-send-email/).
 ---
 
 <table>
    <thead>
        <th>Features</th>
        <th>Currently Available</th>
        <th>Set up Required</th>
    </thead>
    <tbody>
    <tr>
        <td>Register</td>
        <td>V</td>
        <td>X (only store unauthorized user)</td>
    </tr>
    <tr>
        <td>Email Verification</td>
        <td>V</td>
        <td>V</td>
    </tr> 
    <tr>
        <td>Login with JWT</td>
        <td>V</td>
        <td>X(auth.txt file)</td>
    </tr> 
    <tr>
        <td>Restore Password</td>
        <td>X</td>
        <td>V</td>
    </tr> 
    <tr>
        <td>User Setting</td>
        <td>X</td>
        <td>X</td>
    </tr> 
    </tbody>
    
</table>
 
 --- 
 ## Open Source Dependencies

 ### Backend

 > [email-verifier](https://github.com/AfterShip/email-verifier)

 > [Gin Web Framework](https://github.com/gin-gonic/gin)

 > [jwt-go](https://github.com/golang-jwt/jwt)

 > [google/uuid](https://github.com/google/uuid)

 > [Go Cryptography](https://pkg.go.dev/golang.org/x/crypto#section-readme)

 ### Frontend

 > [HTMX](https://github.com/bigskysoftware/htmx)

 > [Bootstrap 5](https://github.com/twbs/bootstrap)

 > [JavaScript Cookie](https://github.com/js-cookie/js-cookie)