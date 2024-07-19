# Service for compiling task lists

## Introduction

The project was created for educational purposes and is an HTTP proxy server written in Golang.  It accepts requests, sends them to the specified external services, and returns a response in JSON format.

## How to install
1. Clone the repository:
   ```sh
   git clone https://github.com/Bekyrys/todo-list
    ```
  ```sh
     cd todo-list
```
2. Build and run using Docker:
```sh
make build
```

3. Run the server:

```sh
make run
```
## Deployed on RENDER:
### Base URL
https://halyk-proxy-server-easy.onrender.com



### Request Body:
POST /api/todo-list/tasks
```sh
{
  "title": "Buy milk"
}
```
### Response example:
```sh
{"id":"0ffcf33a-331e-4f67-be12-cc199490cef1","title":"Buy milk","activeAt":"19 July 17:04","completed":false}
```
**Edit existing tasks:**

PUT /api/todo-list/tasks
```sh
{
  "title": "Buy 2 milk",
  "activeAt": "20 July 13:30"
}
```
### Response example:
```sh
{"id":"0ffcf33a-331e-4f67-be12-cc199490cef1","title":"Buy 2 milk","activeAt":"20 July 13:30","completed":false}
```
**Marks task as done:**

PUT
_/api/todo-list/tasks/{ID}/done_


### Response example:
```sh
{"id":"0ffcf33a-331e-4f67-be12-cc199490cef1","title":"Buy 2 milk","activeAt":"20 July 13:30","completed":true}
```


DELETE /api/todo-list/tasks/{id}
### Response example:

```sh
{"message":"Task deleted"}
```
**Get tasks by status:** 

GET /api/todo-list/tasks?status=active 

**or** 

GET /api/todo-list/tasks?status=done
### Response example:
```sh
[{"id":"0ffcf33a-331e-4f67-be12-cc199490cef1","title":"Buy 2 milk","activeAt":"0000-07-20T13:30:00Z","done":true}]

```
