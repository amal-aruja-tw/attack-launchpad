# attack-launchpad Objective
The objective of this problem is to demonstrate and deep-dive into golang with hands-on experience.<br>
The problem statement can touchbase following aspects when practised under metored environment
1. TDD
2. Golang Basics 
3. Concurrency in Golang (Goroutines and channels) (Not needed)
4. Design Pattern(Strategy Pattern)
5. Databases
6. Web Framework(Golang gin)
7. Docker 
## Problem Premises

### Part 1 - Introduction
It's March end and Ram has to submit his internet expenses for entire year.
He calls his **Shabby** internet provider and requests for his internet bills for this financial year(April till March).<br/>
🕒 <br/>
🕒 <br/>
🕒 <br/>
🕒 <br/>
Time flies!<br/>
It's March 31, and it's only few hours before he can submit his internet bills, and yet there is no word from the **Shabby** Internet Provider. <br/>
Ram being Ram, meanwhile figures out there is an unsecured open endpoint leaking all the user's bills.<br/>
Now ofcourse Ram being non-technical, needs help from someone like you who is an ethical hacker to come to his rescue. The bill will be a CSV file.

#### Supporting contents
* Docker Image of internet provider
* Sample CSV response

### Part 2 - Optimize

The search is taking too long. Is there a way to reduce the time? 
Hint about goroutines

cover - goroutines, wait groups, channels

#### Supporting contents
* https://thoughtworks.udemy.com/course/concurrency-in-go-golang/

### Part 3 - Store

Hearing that Ram was able to get his bill, others have reached out to him for help. Ram decides to store all the bills he was disacrding earlier so that he can share it with anyone who reach out to him.

cover - creation, insertions, migration

Ram realizes he has not been saving field X, which is actually needed for reason Y. So migration is required.

### Part 4 - Share

Ram is receiving a lot of requests for sharing the bills and it has become tedious. He decides to create an API to expose the bill by ID (connection number). 

covers - gin, REST get endpoint

### Part 4.1 - Intro to Middleware

Ram shares a special token to his friends to access his API to avoid unwanted access.

covers - middleware

### Part 5 - Package

Automate db creation, migration and app run

covers - docker, docker-compose, liquibase

### Part 6 - Strategy Pattern

The cable provider has upgraded some users and started giving bill as PDFs. Now, Ram has to decide how to parse the bills based on resposne format.

covers - design pattern - strategy

#### Supporting contents
* Sample PDF response

##  Mentor/Guide Notes
Refer to docker image that spins up a server as **Shabby** Internet Provider.
Extend the problem to cover above 8 points.
The web crawler needs to download all the internet bills as pdfs matching your name. 
If the text is present then save the pdf.

## Context
The Shabby telecom has a range of 712011 entries(7 lakh). We will provide the start and end range.
If it takes even 1 second to process 10 entries, it will take almost 19.778056 hours.
So Linear search approach will take a long time.
Introduce Binary Search
Introduce goroutines
Give importance to times taken to process n entries in each of the above approach, so that the importance and applicability of each can be made understood.




### Building Materials
Contract for Shabby Internet Operator.
Number 1 to 700000

http://localhost:999/internet-bills/:number
```
Content-Type: application/pdf
```

Day1:
Loop over 1 to 7 lakh entries and call endpoint.