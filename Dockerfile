# Specifies a parent image
FROM golang:1.22.4
 
# Creates an app directory to hold your app’s source code
WORKDIR /app/btbschool-service

# Setting Environment Variable
ENV DOMAIN=".aryanattapt.my.id"
ENV PORT="30001"
ENV ENV="PRODUCTION"
ENV POSTGRESQL_HOST="localhost"
ENV POSTGRESQL_PORT="20002"
ENV POSTGRESQL_USERNAME="postgresql"
ENV POSTGRESQL_PASSWORD="postgresql"
ENV MONGODB_URL="mongodb://mongodb:mongodb@localhost:20001"
ENV REDIS_URL=""
ENV REDIS_PASSWORD=""
ENV ASSET_PATH="./assets"
ENV UPLOAD_PATH="./uploads"
ENV JWT_SIGNATURE_KEY="SIGNATURE"

# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
RUN go mod tidy
 
# Builds your app with optional configuration
RUN go build -o btbschool-service

# Tells Docker which network port your container listens on
EXPOSE 30001

# Specifies the executable command that runs when the container starts
#CMD [ “/app/btbschool-service/btbschool-service ]
ENTRYPOINT ["/app/btbschool-service/btbschool-service"]