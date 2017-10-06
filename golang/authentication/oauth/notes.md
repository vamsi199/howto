
details of the oauth url and token url for your program: below repo has details for most of the oauth service providers
- https://github.com/golang/oauth2


where to go to setup the oauth:

**linkedin:**

https://developer.linkedin.com/docs/oauth2


**twitter (WIP):**

*steps per https://dev.twitter.com/web/sign-in/implementing* 


step 1: setup app from https://apps.twitter.com/
fill in id, description, website, callback

step 2: Obtaining a request token: POST oauth / request_token

step 3: Redirecting the user: GET oauth / authenticate


*errors list:*

1) 

Whoa there!
There is no request token for this page. That's the special key we need from applications asking to use your Twitter account. Please go back to the site or application that sent you here and try again; it was probably just a mistake.


*links:*

http://pierrecaserta.com/go-oauth-facebook-github-twitter-google-plus/

https://github.com/mrjones/oauth/blob/master/examples/twitter/twitter.go

Sign In With Twitter: Create A Twitter App (2/5) https://www.youtube.com/watch?v=E0SmdBobo9Y

Sign In With Twitter: Authenticating (4/5) https://www.youtube.com/watch?v=ZOej-9JUzyQ

Implementing Sign in with Twitter: https://dev.twitter.com/web/sign-in/implementing

GET oauth/authenticate https://developer.twitter.com/en/docs/basics/authentication/api-reference/authenticate

POST oauth/request_token https://developer.twitter.com/en/docs/basics/authentication/api-reference/request_token

