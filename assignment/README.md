## GlobalWebIndex Engineering Challenge 

### The intention of the given implementation was to apply and present good programming practices. There are sections of the implementation that due to time restrictions were not fully completed but i hope that the main idea and the intention of each decision is distinct and clear.

### Architectual Decisions

The implementation divides the server application to:
- data access layer communicating with the data base
- service layer including the business logic
- middleware layer for authorization purposes
- controllers layer for routing the requests to the equivalent endpoints

The aforementioned decisions took into consederation scalability issues, authorization patterns and separation of concerns

A completed testing implementation was not fullfilled

The approach attempts to even briefly display clean code and architecture practices, dependency injection usage, cors handling, naming conventions assisting readability and nice to implement coding practices in general

I would like to note that there are unnecessary comments that in a different concept would be cleared but their purpose is mainly to refer to uncompleted parts of the implementation, sorry for that

##### Api Endpoints:

###### Authorize
authorize the access

###### Asset
handle asset requests 
###### Chart/Audience/Insights
handle specified asset requests, possible future endpoints likely added here regarding specific assets

##### Libraries/Tools Utilized:

- gorm for ORM https://github.com/go-gorm/gorm 
- envconfig for configurating environment variables https://github.com/kelseyhightower/envconfig
- yaml to encode/decode yaml files
- gin Web framework https://github.com/gin-gonic/gin
- Docker for source containerization https://github.com/docker
- jwt tokes for authentication https://github.com/auth0/node-jsonwebtoken