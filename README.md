# GO-Register

### Register/Login web app, Restiful APIs and JWT 

 ---
 > Go-Register is a full-stack, single-page web app provides features around user register and login.

 > It focuses on user experience, extensibility, and performance.

 > If you are...

 * A web dev who wants to learn the user register feature.
 * A designer who wants to make UI/UX around user login.
 * Or doesn't want to make the user login feature from scratch :P
 
 > Then this is for you! 
 ---
 ## Before you start...
 1. Data storage
   >> User Data are stored in hashmap, thus the app can start without the database. But keep in mind using a database to persist data is very important. 
 2. JWT generate
   >> JWT are produced by the key in auth.txt, if you want to use this project elsewhere, remember to change the key to prevent security vulnerability. 
 ---
## How to start

 > clone this project:
    
    git clone https://github.com/guava-coder/go-register.git

 > If you have golang install, go into project folder, and type inputs below in terminal:
    
    go run .

 ---
 
 <table>
    <thead>
        <th>Features</th>
        <th>Currently Available</th>
        <th>Require set up</th>
    </thead>
    <tbody>
    <tr>
        <td>Register</td>
        <td>V</td>
        <td>X</td>
    </tr>
    <tr>
        <td>Email Verification</td>
        <td>V</td>
        <td>V</td>
    </tr> 
    <tr>
        <td>Login with JWT</td>
        <td>V</td>
        <td>X</td>
    </tr> 
    <tr>
        <td>Restore Password</td>
        <td>X</td>
        <td>V</td>
    </tr> 
    <tr>
        <td>User Setting Display</td>
        <td>X</td>
        <td>X</td>
    </tr> 
    <tr>
        <td>Change Email</td>
        <td>X</td>
        <td>V</td>
    </tr> 
    <tr>
        <td>Change Password</td>
        <td>X</td>
        <td>V</td>
    </tr> 
    </tbody>
    
</table>