# Start from golang v1.14 base image
FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR $/Users/Negar/Desktop/vscode/Usage

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN GO111MODULE=on

 EXPOSE 80

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["usage", "run", "--host", "0.0.0.0"]
