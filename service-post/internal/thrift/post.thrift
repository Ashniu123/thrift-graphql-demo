namespace go post
namespace js post

typedef i32 int

struct PingResponse {
  1: optional string message
  2: optional int version
}

service PostService {
  PingResponse ping(),
}
