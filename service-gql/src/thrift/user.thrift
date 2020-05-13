namespace go user
namespace js user

typedef i32 int

struct PingResponse {
  1: optional string message
  2: optional int version
}

service UserService {
  PingResponse ping(),
}
