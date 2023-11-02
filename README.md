# grpc-user

## Installation

1. **Go**: Make sure you have Go 1.16 or later installed. 

2. **Docker**: Install Docker if you don't have it. 

3. **Clone Repository**: Clone this repository.

```bash
git clone https://github.com/tanishqtqs/grpc-user.git
cd grpc-user
docker build -t grpc-user .
docker run -p 50051:50051 grpc-user
