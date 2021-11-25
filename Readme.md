# Setting up Environment
In order to use DBMigration repository, you need to
* Install [Go](https://golang.org/dl/) on your local system
* Add /usr/local/go/bin to the PATH environment variable.
    * You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

          export PATH=$PATH:/usr/local/go/bin

    * Next, run the shell commands directly or execute them from the profile using a command such as

          source $HOME/.profile

* Install Docker on local system, and pull mongo image and run the container

# COVID-19 Impact Analysis API:

Johns Hopkins University Center for Systems Science and Engineering (JHU CSSE) publishes
data about Coronavirus COVID-19 impact on a daily basis.

Use the following command to create a mongo container on your system.

    docker run -p 27017:27017 --name user-mongo-db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=TestPass1234 -d mongo:4.2.0

Now to start the controller use the following command:

     go run cmd/impact-analysis-server/main.go

Architecture
    
    Microservice architecture is divided into following layers:
    1- API Layer handles API payloads and responses.
    2- Service Layer contains the whole business ligic.
    3- DB Layer handles all database transactions

Assumption

    This service is speciically written for the given dataset.

API Doc
    
    Swagger Documentation for API's
    https://app.swaggerhub.com/apis/abdulmoeid7112/ImpactAnalysisAPI/1.0.0