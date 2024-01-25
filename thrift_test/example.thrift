namespace go hello.cloudwego.team

struct MyReq{
    1:required string name,
    2:required string id
}

struct MyResp{
    1:required string text
}

service greet {
	MyResp Hello(1:required MyReq req)
}

