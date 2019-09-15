####################################
# First stage: build the executable.
####################################
FROM golang:alpine AS builder

# Install Git to allow go mod to download dependencies.
RUN apk update && apk add --no-cache git

# Create the app user.
RUN adduser -D -g '' appuser

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Download and verify dependencies as a separate layer.
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

# Copy the source code and build the executable.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

###########################################
# Final stage: the runtime container image.
###########################################
FROM scratch AS final

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

# Copy the executable from the build stage.
COPY --from=builder /app/skills-service /app/

# Set the app listen port.
ENV APP_ADDR ":50002"

# Expose the app listen port.
EXPOSE 50002

# Perform any further action as an unprivileged user.
USER appuser

# Run the compiled binary.
ENTRYPOINT ["/app/skills-service"]
