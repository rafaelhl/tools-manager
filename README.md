## Tools Manager API

This project implements a tools manager API with these three operations:

 - GET /tools?tag={tag_filter} (Fetch all tools being possible filtering by a tag)
 - POST /tools (Create a new tool)
 - DELETE /tools/:id (Delete a tool by ID)
 
You can see how to use these operations following the yaml of Swagger. [click here](resources/specs/swagger.yaml)

## How to run

To up the application only use the following Go command:

```shell script
    go run .
```
  
To validate if all up ok send an HTTP GET to `/ping` route and expect a text `pong`:

```shell script
    curl http://localhost:3000/ping   
```

