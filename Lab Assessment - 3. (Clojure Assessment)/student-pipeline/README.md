# Student Performance Data Pipeline

## Student Details

Name: Aditya Chouhan  
Roll No: 2401830001  
Course: B.Sc. (H) CyberSecurity  
Semester: V

## Project Description

An integrated Clojure application that processes student records using map, filter and reduce functions, stores processed results in MongoDB and exposes data through Ring HTTP endpoints.

## Technologies Used

- Clojure
- Java 25
- Leiningen
- MongoDB
- Ring HTTP
- Jetty
- Cheshire JSON

## Execution Instructions

1. Start the MongoDB service.
2. Open the student-pipeline project folder in a terminal.
3. Execute `lein deps`.
4. Execute `lein run`.
5. Open http://localhost:3000 in a web browser.

## API Endpoints

- `/` — Application homepage
- `/students` — All student records
- `/passed` — Passed students
- `/failed` — Failed students
- `/statistics` — Performance statistics

## Results

- Total students: 6
- Passed students: 5
- Failed students: 1
- Average marks: 66.17

## Conclusion

The application successfully demonstrates functional programming, database integration, exception handling, Java interoperability and web API development using Clojure.
