import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

// Import service definition that you want to connect to.
import { GreetService } from "./protogen/greet/v1/greet_connect";
import { GreetRequest } from "./protogen/greet/v1/greet_pb";

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
// If your endpoint only supports gRPC-web, make sure to use
// `createGrpcWebTransport` instead.
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(GreetService, transport);

function App() {
  const promiseGreet = client.greet({ name: "bow wow"})
  promiseGreet.then((res) => {
    console.log(res.HelloOneOf)
    console.log(res.helloAny?.is(GreetRequest))
  })
  return <>Hello world</>;
}

export default App
