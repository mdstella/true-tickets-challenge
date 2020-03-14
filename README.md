# true-tickets-challenge

### Changelog and decisions

1. The project structure was done at the same time with the AddMetric endpoint definition, to use it as an example and implementation of the structure. The structure has the following folders
   1. dao --> will keep all the persistence logic, for the challenge purpose all the persisten will be done in memory instead of using an storage
   2. endpoint --> this folder is the entry point for the invocations, here I defined the routing and endpoints.
   3. errors --> it has the errors definitions for the API
   4. model --> in this folder I will define the data structures that will use for the API (Request/Response objects and DTOs)
   5. service --> here I will put all the business logic and interactions with the persistence layer.
2. Addnig the /metric/{key} endpoint.
   1. Create a new endpoint: `POST --> /metric/{key}` that uses a in memory dao to persist/loads the metric value based on the given key.
   2. The DAO layer is defined by an interface, as I already mentioned the actual implementation is an in memory storage, but with the interface definitions it's really simple to extend a new implmenentation to use MongoDB, MySQL, Redis or the persistance tool that you need.
3. Adding the /metric/{key}/sum endpoint. This will retrieve the metric key sum
   1. Added the new endpoint `GET --> /metric/{key}/sum`.
   2. On the DAO layer we are retrieving the metrics that aren't expired. I'm considering the 1 hour TTL by default. If I find an expired metric I'm not retrieving it and also removing it from the memory map.
4. Adding API documentation using Swagger
   1. There is a new endpoint expose that is `http://localhost:9091/swagger-ui/swagger.json`. Invoking it will retrieve the Swagger documentation JSON.
   2. You can download the Chrome extension [Swagger UI console](https://chrome.google.com/webstore/detail/swagger-ui-console/ljlmonadebogfjabhkppkoohjkjclfai?utm_source=chrome-ntp-icon) and add that endpoint on the extension, then click `Explore` and will provide a UI to see the documentation and interact with the API (To do this the server must be up and running)
   3. I also added to interact/documentation a Postman collection on the folder `/postman` on the repository.
5. Adding `Dockerfile` to be able to run the application using docker to avoid installing GO on your computer
   1. To build the Docker image I just go to `true-tickets-challenge` folder and run: `docker build -t true-tickets-challenge .` (maybe `sudo` is needed to execute the command). That will generate the image, first time can take some time.
   2. To run the image I run: `docker run -p 9091:9091 true-tickets-challenge`. This will startup the application and expose it on the port 9091.
   3. If you want you can assign a name to the image (`docker run -p 9091:9091 --name=TTC true-tickets-challenge`), but then you will need to remove the container once you stop the docker image. To do that:
      1. Run `docker ps -a` to see the already used containers.
      2. Run `docker rm TTC` (TTC or the name you used for the container).
      3. After that you will be able to start it again