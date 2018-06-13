/**
install npm package
    npm install google-protobuf --save
    npm install grpc --save
#run go service
$ go run inf_server.go
service has run in  :50051
request data: name:"daheige" ua:2

$ node hello_client.js
*/
var grpc = require('grpc');
const PROTO_PATH = './proto/inf.proto';

const port = '50051';
const host = "127.0.0.1";

const proto = grpc.load(PROTO_PATH).helloworld;
// console.log(proto)

//Greeter service name
const client = new proto.Greeter(host + ":" + port, grpc.credentials.createInsecure());
// console.log(client)
// console.log(typeof proto.HelloRequest);

//执行rpc调用
// var request = new proto.HelloRequest();
// request.setName("daheige")
// request.setUa(2)

//method1:
/*obj as params
client.SayHello({
    name: "daheige",
    ua: 2,
}, function(err, data) {
    if (err) console.log(err)
    console.log("reply data : ", data)
    console.log(data.message)
    grpc.closeClient(client);
});
*/

//method2:
client.SayHello({
    name: "daheige",
    ua: 2,
}, function(error, data) {
    if (error) {
        console.log(error)
    }
    console.log(data)
});
