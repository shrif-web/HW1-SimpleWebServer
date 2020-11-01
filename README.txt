Welcome to our simple webserver!

This is the first assignment of web-programming course, fall semester, 2020.

Group members are:
- Ali Jandaghi Alaee
- Sepehr Ashrafzadeh
_ Ashkan Larni

This project is designed for two perpuses,
1 - Calculation of summation of two input integers and its encoding using sha256.
2 - Showing the requested line of a sample .txt file existing in server.

Frontend of this project is written in html and js, and the backend is handled by two servers one written in node js and the other one in go.
This project implemented and tested on a CentOS system managed by nginx.

-------

Here is the results of request analysis of our two servers using Locust;


 ————— Write API —————-
 - go handled ~170 req/sec by 2.5% of failure
- nodes handles ~330 req/sec by 0% of failure


 ————— SHA256 API —————-
 - go handled ~165 req/sec by 1% of failure
- nodes handles ~333 req/sec by 0% of failure

